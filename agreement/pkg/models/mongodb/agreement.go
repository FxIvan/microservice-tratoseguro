package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/fxivan/microservicio/agreement/pkg/models"
	"github.com/go-zoox/fetch"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AgreementModel struct {
	C *mongo.Collection
}

type SearchInfoBody struct {
	Username string
	Email    string
}

const charsetAlphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetNumeric = "0123456789"

func GenerateUUID(length int, typeOfKey string) string {
	rand.Seed(time.Now().UnixMicro())

	var charset string

	switch typeOfKey {
	case "numeric":
		charset = charsetNumeric
	case "alphanumeric":
		charset = charsetAlphanumeric
	}

	idRandom := make([]byte, length)

	for i := range idRandom {
		idRandom[i] = charset[rand.Intn(len(charset))]
	}
	return string(idRandom)
}

func (m AgreementModel) SaveAgreement(contract *models.ContractDefinitionModel, emailUser string) (string, bool) {

	searchUserCTPY := SearchInfoBody{
		Username: contract.Counterparty.Username,
		Email:    contract.Counterparty.EmailCTPY,
	}

	searchUserPRNE := SearchInfoBody{
		Email: emailUser,
	}

	responseCTPY, err := fetch.Post("http://host.docker.internal:9090/auth/info", &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: searchUserCTPY,
	})

	reponsePRNE, err := fetch.Post("http://host.docker.internal:9090/auth/info", &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: searchUserPRNE,
	})

	if err != nil {
		return "Error al buscar el usuario", false
	}

	responseJSONPRNE, err := reponsePRNE.JSON()
	responseJSONCTPY, err := responseCTPY.JSON()

	if err != nil {
		return "Error al buscar el usuario", false
	}

	bytesCTPY := []byte(responseJSONCTPY)
	bytesPRNE := []byte(responseJSONPRNE)

	type Resultados struct {
		Username string
		Email    string
		ID       string
	}

	var resultadosCTPY Resultados
	json.Unmarshal(bytesCTPY, &resultadosCTPY)

	var resultadosPRNE Resultados
	json.Unmarshal(bytesPRNE, &resultadosPRNE)

	idRandom := GenerateUUID(16, "alphanumeric")

	_, err = m.C.InsertOne(context.TODO(), bson.M{
		"contractIdentifier": idRandom,
		"counterparty": bson.M{
			"idRefCTPY": contract.Counterparty.IDRefCTPY,
			"username":  resultadosCTPY.Username,
			"emailCTPY": resultadosCTPY.Email,
			"iduser":    resultadosCTPY.ID,
		},
		"proponent": bson.M{
			"emailPRNE":    resultadosPRNE.Email,
			"usernamePRNE": resultadosPRNE.Username,
			"iduserPRNE":   resultadosPRNE.ID,
		},
		"agreementText": bson.M{
			"text": contract.AgreementText.Text,
		},
		"counterparty_signature": bson.M{
			"dni":      "",
			"fullName": "",
			"accepte":  "",
		},
		"proposing_firm": bson.M{
			"dni":      contract.ProposingFirm.DNI,
			"fullName": contract.ProposingFirm.FullName,
			"accepte":  contract.ProposingFirm.Accepte,
		},
		"agreement_status": bson.M{
			"status": "Created",
			"text":   "Contrato creado por el proponente",
		},
		"add_field_required": contract.AddFieldRequired,
	})

	if err != nil {
		fmt.Print(err)
		return "Error al crear el contrato", false
	}

	return "Contrato creado correctamente", true
}

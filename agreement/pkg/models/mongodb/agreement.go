package mongodb

import (
	"context"
	"encoding/json"
	"fmt"

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

	_, err = m.C.InsertOne(context.TODO(), bson.M{
		"counterparty": bson.M{
			"idRefCTPY":     contract.Counterparty.IDRefCTPY,
			"username":      resultadosCTPY.Username,
			"emailCTPY":     resultadosCTPY.Email,
			"iduser":        resultadosCTPY.ID,
			"linkShareCTPY": contract.Counterparty.LinkShareCTPY,
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
			"dni":      contract.CounterpartySignature.DNI,
			"fullName": contract.CounterpartySignature.FullName,
			"accepte":  contract.CounterpartySignature.Accepte,
		},
		"proposing_firm": bson.M{
			"dni":      contract.ProposingFirm.DNI,
			"fullName": contract.ProposingFirm.FullName,
			"accepte":  contract.ProposingFirm.Accepte,
		},
		"agreement_status": bson.M{
			"status": contract.AgreementStatus.Status,
			"text":   contract.AgreementStatus.Text,
		},
		"contractLinkId": bson.M{
			"id":       contract.ContractLinkId.ID,
			"password": contract.ContractLinkId.Password,
		},
		"add_field_required": contract.AddFieldRequired,
	})

	if err != nil {
		fmt.Print(err)
		return "Error al crear el contrato", false
	}

	return "Contrato creado correctamente", true
}

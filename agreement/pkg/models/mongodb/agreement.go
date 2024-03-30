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

func (m AgreementModel) SaveAgreement(contract *models.ContractDefinitionModel) (string, bool) {

	searchUser := SearchInfoBody{
		Username: contract.Counterparty.Username,
		Email:    contract.Counterparty.EmailCTPY,
	}

	response, err := fetch.Post("http://host.docker.internal:9090/auth/info", &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: searchUser,
	})

	if err != nil {
		return "Error al buscar el usuario", false
	}

	responseJSON, err := response.JSON()

	if err != nil {
		return "Error al buscar el usuario", false
	}

	bytes := []byte(responseJSON)

	type Resultados struct {
		Username string
		Email    string
		ID       string
	}

	var resultados Resultados
	json.Unmarshal(bytes, &resultados)

	fmt.Println(resultados)
	fmt.Println(resultados.ID)

	_, err = m.C.InsertOne(context.TODO(), bson.M{
		"counterparty": bson.M{
			"idRefCTPY":     contract.Counterparty.IDRefCTPY,
			"username":      resultados.Username,
			"emailCTPY":     resultados.Email,
			"iduser":        resultados.ID,
			"linkShareCTPY": contract.Counterparty.LinkShareCTPY,
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

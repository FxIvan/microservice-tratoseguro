package mongodb

import (
	"context"
	"fmt"

	"github.com/fxivan/microservicio/agreement/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AgreementModel struct {
	C *mongo.Collection
}

func (m AgreementModel) SaveAgreement(contract *models.ContractDefinitionModel) (string, bool) {

	fmt.Print(contract)

	_, err := m.C.InsertOne(context.TODO(), bson.M{
		"counterparty": bson.M{
			"idRefCTPY":     contract.Counterparty.IDRefCTPY,
			"emailCTPY":     contract.Counterparty.EmailCTPY,
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

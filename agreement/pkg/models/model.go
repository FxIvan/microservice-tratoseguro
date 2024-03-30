package models

type NameFileCTPYObject struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type StatusLoadObject struct {
	Name string `json:"name"`
	Load bool   `json:"load"`
}

type FieldPRNObject struct {
	FieldName string `json:"fieldName"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type CounterpartFilenameModel struct {
	NameFileCTPY    NameFileCTPYObject `json:"nameFileCTPY"`
	DniFrontCTPY    StatusLoadObject   `json:"dniFrontCTPY"`
	DniBackCTPY     StatusLoadObject   `json:"dniBackCTPY"`
	LocationService StatusLoadObject   `json:"locationService"`
	Email           string             `json:"email"`
	FacePhoto       StatusLoadObject   `json:"facePhoto"`
	FieldOnePRNE    FieldPRNObject     `json:"fieldOnePRNE"`
	FieldTwoPRNE    FieldPRNObject     `json:"fieldTwoPRNE"`
	IdUserCTPY      string             `bson:"_id"`
	IdUserPRNE      string             `bson:"_id"`
}

type CounterpartyObject struct {
	IDRefCTPY     string `json:"idRefCTPY"`
	EmailCTPY     string `json:"emailCTPY"`
	Username      string `json:"username"`
	IDUser        string `json:"iduser"`
	LinkShareCTPY string `json:"linkShareCTPY"`
}

type AgreementTextObject struct {
	Text string `json:"text"`
}

type CounterpartySignatureObject struct {
	DNI      string `json:"dni"`
	FullName string `json:"fullName"`
	Accepte  bool   `json:"accepte"`
}

type ProposingFirmObject struct {
	DNI      string `json:"dni"`
	FullName string `json:"fullName"`
	Accepte  bool   `json:"accepte"`
}

type AgreementStatusObject struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}

type ContractLinkIdObject struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type AddFieldRequiredArray struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

type ContractDefinitionModel struct {
	Counterparty          CounterpartyObject          `json:"counterparty"`
	AgreementText         AgreementTextObject         `json:"agreementText"`
	CounterpartySignature CounterpartySignatureObject `json:"counterparty_signature"`
	ProposingFirm         ProposingFirmObject         `json:"proposing_firm"`
	AgreementStatus       AgreementStatusObject       `json:"agreement_status"`
	ContractLinkId        ContractLinkIdObject        `json:"contractLinkId"`
	AddFieldRequired      []AddFieldRequiredArray     `json:"add_field_required"`
}

type SearchUser struct {
	IdUser string `json:"iduser"`
	Email  string `json:"email"`
}

type PhotoObject struct {
	Path string `json:"path"`
}

type UserSignup struct {
	ID         string      `bson:"_id"`
	Username   string      `json:username`
	Password   string      `json:password`
	Email      string      `json:email`
	Phone      string      `json:phone`
	Name       string      `json:name`
	LastName   string      `json:lastName`
	Address    string      `json:address`
	City       string      `json:city`
	Country    string      `json:country`
	PostalCode string      `json:postalCode`
	Building   string      `json:building`
	Apartment  string      `json:apartment`
	Active     bool        `json:active`
	DNI        string      `json:dni`
	Gender     string      `json:gender`
	PhotoFront PhotoObject `json:photoFront`
	PhotoBack  PhotoObject `json:photoBack`
	FacePhoto  PhotoObject `json:facePhoto`
}

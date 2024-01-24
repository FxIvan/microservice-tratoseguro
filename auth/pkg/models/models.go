package models

type UserLogin struct {
	Username string `json:username`
	Password string `json:password`
}

type UserSignup struct {
	Username        string `json:username`
	Password        string `json:password`
	ConfirmPassword string `json:confirmPassword`
	Email           string `json:email`
	Phone           string `json:phone`
	Name            string `json:name`
	LastName        string `json:lastName`
	Address         string `json:address`
	City            string `json:city`
	Country         string `json:country`
	PostalCode      string `json:postalCode`
	Building        string `json:building`
	Apartment       string `json:apartment`
}

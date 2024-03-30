package models

type UserLogin struct {
	Username string `json:username`
	Password string `json:password`
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

type RequestInfoUser struct {
	Email    string `json:"email"`
	ID       string `bson:"_id"`
	Username string `json:"username"`
}

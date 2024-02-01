package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/auth/pkg/models"
	"github.com/fxivan/microservicio/auth/pkg/sendgrid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserSignupModel struct {
	C *mongo.Collection
}

func HashBcrypt(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), 10)
	return string(hash), err
}

func (m UserSignupModel) RegisterUser(userSignup *models.UserSignup) (string, bool) {

	encryptedPassword, err := HashBcrypt(userSignup.Password)
	if err != nil {
		return "Error en la encriptacion de la contraseña", false
	}

	filterUsername := bson.M{"username": userSignup.Username}

	findUsername := m.C.FindOne(context.Background(), filterUsername)

	if findUsername.Err() == nil {
		return "El usuario ya existe", false
	}

	filterEmail := bson.M{"email": userSignup.Email}

	findEmail := m.C.FindOne(context.Background(), filterEmail)
	if findEmail.Err() == nil {
		return "El email ya existe", false
	}

	sendgrid.SendEmailSengrid(userSignup.Email, "Este es un mensaje de testeo")

	_, err = m.C.InsertOne(context.TODO(), bson.M{
		"username":   userSignup.Username,
		"password":   encryptedPassword,
		"email":      userSignup.Email,
		"phone":      userSignup.Phone,
		"name":       userSignup.Name,
		"lastName":   userSignup.LastName,
		"address":    userSignup.Address,
		"city":       userSignup.City,
		"country":    userSignup.Country,
		"postalCode": userSignup.PostalCode,
		"building":   userSignup.Building,
		"apartment":  userSignup.Apartment,
		"active":     userSignup.Active,
	})
	if err != nil {
		return "Error al crear al usuario", false
	}

	return "Usuario creado correctamente", true
}

func (m UserSignupModel) FindUsername(username string) (*models.UserSignup, error) {
	var user models.UserSignup

	filter := bson.M{"username": username}
	err := m.C.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

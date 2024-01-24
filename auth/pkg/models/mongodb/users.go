package mongodb

import (
	"context"
	"fmt"

	"github.com/fxivan/microservicio/auth/pkg/models"
	"github.com/fxivan/microservicio/auth/pkg/response"
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

func (m UserSignupModel) InsertRegisterUser(userSignup *models.UserSignup) (*response.Response, error) {

	fmt.Println("Insertando Usuario", userSignup)

	encryptedPassword, err := HashBcrypt(userSignup.Password)

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
		return nil, err
	}

	response := &response.Response{
		Status:  "success",
		Message: "Created user .",
		Code:    200,
	}

	return response, nil
}

func (m UserSignupModel) FindUserEmail(username string) (*models.UserSignup, error) {
	var user models.UserSignup
	filter := bson.M{"username": username}
	err := m.C.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

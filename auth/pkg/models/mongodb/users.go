package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/auth/pkg/models"
	"github.com/fxivan/microservicio/auth/pkg/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSignupModel struct {
	C *mongo.Collection
}

func (m UserSignupModel) InsertRegisterUser(userSignup *models.UserSignup) (*response.Response, error) {
	_, err := m.C.InsertOne(context.TODO(), userSignup)
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

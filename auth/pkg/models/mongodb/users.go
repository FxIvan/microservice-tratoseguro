package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/auth/pkg/models"
	"github.com/fxivan/microservicio/auth/pkg/response"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSignInModel struct{
		C *mongo.Collection
}

func (m UserSignInModel) InsertRegisterUser(userSignIn *models.UserSingIn) (*response.Response, error){
	_ ,err := m.C.InsertOne(context.TODO(),userSignIn)
	if err != nil {
		return nil,err
	}

	response := &response.Response{
		Status:  "success",
		Message: "Created user .",
		Code:    200,
	}

	return response,nil
}
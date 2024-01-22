package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/auth/pkg/models"
	"github.com/fxivan/microservicio/auth/pkg/response"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSignupModel struct{
		C *mongo.Collection
}

func (m UserSignupModel) InsertRegisterUser(userSignup *models.UserSignup) (*response.Response, error){
	_ ,err := m.C.InsertOne(context.TODO(),userSignup)
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
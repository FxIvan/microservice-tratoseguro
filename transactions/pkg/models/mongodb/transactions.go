package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/transactions/pkg/models"
	"github.com/fxivan/microservicio/transactions/pkg/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionModel struct {
	C *mongo.Collection
}


func (m *TransactionModel) InsertTransaction(transaction *models.Transaction) (*response.Response,error) {
	result , err := m.C.InsertOne(context.TODO(), transaction)
	if err != nil {
		return nil, err
	}

	insertResult := &response.Response{
		Status:  "success",
		Message: "Transaction inserted successfully with ID " + result.InsertedID.(primitive.ObjectID).Hex() + ".",
		Code:    200,
	}

	return insertResult, nil
}
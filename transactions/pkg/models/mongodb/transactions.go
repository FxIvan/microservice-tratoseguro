package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/transactions/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionModel struct {
	C *mongo.Collection
}


func (m *TransactionModel) InsertTransaction(transaction *models.Transaction) (*mongo.InsertOneModel,error) {
	_, err := m.C.InsertOne(context.TODO(), transaction)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
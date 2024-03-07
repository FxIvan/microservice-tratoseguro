package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/agreement/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSignupModel struct {
	C *mongo.Collection
}

func (m UserSignupModel) getAllUsers() ([]*models.UserSignup, error) {
	var users []*models.UserSignup
	cursor, err := m.C.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.UserSignup
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

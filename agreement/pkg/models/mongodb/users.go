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

func (m UserSignupModel) searchUser(UserSignup *models.UserSignup) (string, bool) {

	filter := bson.M{"username": UserSignup.Username}

	findUsername := m.C.FindOne(context.Background(), filter)

	if findUsername.Err() != nil {
		return "El usuario no existe", false
	}

	return "Usuario existe", true
}

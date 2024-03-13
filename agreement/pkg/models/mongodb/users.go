package mongodb

import (
	"context"
	"fmt"

	"github.com/fxivan/microservicio/agreement/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSignupModel struct {
	C *mongo.Collection
}

func (m UserSignupModel) SearchUser(UserSignup *models.SearchUser) (string, bool) {

	filter := bson.M{"email": UserSignup.Email}

	cursor := m.C.FindOne(context.Background(), filter)

	if cursor.Err() != nil {
		fmt.Println("Usuario no existe")
		return "El usuario no existe", false
	}

	return "Usuario existe", true

}

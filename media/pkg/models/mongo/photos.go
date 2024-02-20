package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/media/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhotosModel struct {
	C *mongo.Collection
}

func (m PhotosModel) UploadImage(photoModel *models.ModelPhoto, idUser string) (string, bool) {
	_, err := m.C.InsertOne(context.TODO(), bson.M{
		"userId":    photoModel.UserId,
		"nameImg":   photoModel.NameImg,
		"size":      photoModel.Size,
		"createdAt": photoModel.CreatedAt,
	})
	if err != nil {
		return "Error al crear el archivo", false
	}

	return "Foto Subida", true
}

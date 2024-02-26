package mongodb

import (
	"context"

	"github.com/fxivan/microservicio/media/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilesModel struct {
	C *mongo.Collection
}

func (m FilesModel) AddFile(fileModel *models.ModelFile) (string, bool) {
	_, err := m.C.InsertOne(context.TODO(), bson.M{
		"userId":    fileModel.UserId,
		"contract":  fileModel.NameFile,
		"size":      fileModel.Size,
		"createdAt": fileModel.CreatedAt,
	})

	if err != nil {
		return "Error al crear el archivo", false
	}

	return "Contrato Subido", true
}

func (m FilesModel) FindFile(nameFile string) (*models.ModelFile, bool) {

	var fileModel models.ModelFile

	filter := bson.M{"nameImg": nameFile}
	err := m.C.FindOne(context.TODO(), filter).Decode(&fileModel)
	if err != nil {
		return nil, false
	}

	return &fileModel, true
}

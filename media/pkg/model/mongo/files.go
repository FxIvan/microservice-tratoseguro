package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type FilesModel struct {
	C *mongo.Collection
}

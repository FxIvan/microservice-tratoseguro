package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type PhotosModel struct {
	C *mongo.Collection
}

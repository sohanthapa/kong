package mongostore

import (
	"context"
	"kong/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateVersion(collection *mongo.Collection, version *models.Version) error {
	_, err := collection.InsertOne(context.Background(), &version)
	if err != nil {
		return err
	}
	return nil

}

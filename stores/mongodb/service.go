package mongostore

import (
	"context"
	"kong/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateService(collection *mongo.Collection, service *models.Service) error {
	_, err := collection.InsertOne(context.Background(), &service)
	if err != nil {
		return err
	}
	return nil
}

func ListServices() {

}

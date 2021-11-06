package mongostore

import (
	"context"
	"kong/models"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CreateService inserts a new service into the service collection
func CreateService(collection *mongo.Collection, service *models.Service) error {
	_, err := collection.InsertOne(context.Background(), &service)
	if err != nil {
		return err
	}
	return nil
}

//GetService finds and returns the service based on the serviceID
func GetService(db *mongo.Database, serviceID string) (models.Service, error) {
	serviceCollection := db.Collection("services")
	var service models.Service
	err := serviceCollection.FindOne(context.Background(), bson.M{"id": serviceID}).Decode(&service)
	if err != nil {
		return service, err
	}
	return service, nil
}

//ListServices lists all the services using the offset, limit and filter parameter
func ListServices(db *mongo.Database, slQuery models.ServiceListQuery) ([]models.Service, error) {
	serviceCollection := db.Collection("services")
	findOptions := options.Find()
	findOptions.SetSkip(slQuery.Offset)
	findOptions.SetLimit(slQuery.Limit)
	if slQuery.Sort == "desc" {
		findOptions.SetSort(bson.D{{"name", -1}})
	}
	cursor, err := serviceCollection.Find(context.Background(), slQuery.Filter, findOptions)
	if err != nil {
		return nil, err
	}

	services := make([]models.Service, 0)
	err = cursor.All(context.Background(), &services)
	if err != nil {
		return nil, err
	}
	return services, nil

}

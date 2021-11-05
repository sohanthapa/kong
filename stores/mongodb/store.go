package mongostore

import (
	"context"
	"log"
	"time"

	"kong/data"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://testdb:test@cluster0.zdhnz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func Initialize(client *mongo.Client) {
	db := client.Database("testdb")
	serviceCollection := db.Collection("services")
	versionCollection := db.Collection("versions")

	//populate versions into the database
	for _, version := range data.VersionList {
		err := CreateVersion(versionCollection, &version)
		// log the error but continue to insert other version from the list
		if err != nil {
			log.Printf("fail to insert version %s", version.ID)
			continue
		}
	}

	//populate services into the database
	for _, service := range data.ServiceList {
		err := CreateService(serviceCollection, &service)
		// log the error but continue to insert other service from the list
		if err != nil {
			log.Printf("fail to insert service %s", service.ID)
			continue
		}
	}

}

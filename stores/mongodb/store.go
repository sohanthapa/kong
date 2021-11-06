package mongostore

import (
	"context"
	"fmt"
	"log"
	"time"

	"kong/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//New initializes a new mongo db based on the URI and database name
// NOTE: for the simplicity of this exercise I choose to hardcode the URI and database name
func New() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb+srv://testdb:test@cluster0.zdhnz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("testdb")
	return db
}

// Populate function populates the data (services, versions) into the db
func Populate(db *mongo.Database) {
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

// CreateIndex - creates an index for a specific field in a collection
func CreateIndex(db *mongo.Database, fieldName, collection string) bool {

	// 1. Lets define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    bson.M{fieldName: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(true),
	}

	// 2. Create the context for this operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Connect to the database and access the collection
	col := db.Collection(collection)

	// 4. Create a single index
	_, err := col.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

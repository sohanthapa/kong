package main

import (
	"context"
	"fmt"
	"kong/models"
	"kong/server"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb+srv://testdb:test@cluster0.zdhnz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("\n\n hola error is %v\n\n", err)
		return
	}

	db := client.Database("testdb")
	serviceCollection := db.Collection("services")
	s := models.Service{}
	_, err = serviceCollection.InsertOne(context.Background(), &s)
	if err != nil {
		log.Println(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	var sCfg server.Config

	server := server.New(&sCfg)
	server.Initialize()

	httpServer := &http.Server{
		//Addr:    fmt.Sprintf(":%d", defaultPort),
		Handler: server.Router,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

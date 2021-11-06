package main

import (
	"fmt"
	"kong/server"
	"log"
	"net/http"

	mongostore "kong/stores/mongodb"

	"github.com/go-chi/chi"
)

const (
	// default settings
	defaultPort int = 8080
)

func main() {

	//call to initialize mongodb connnection client
	db := mongostore.New()

	//populate the data in db
	mongostore.Populate(db)

	//create index for both the collections to use id as unique value
	mongostore.CreateIndex(db, "id", "services")
	mongostore.CreateIndex(db, "id", "versions")

	server := &server.Server{
		Router:     chi.NewRouter(),
		MongoStore: db,
	}

	//call to initialize server
	server.Initialize()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", defaultPort),
		Handler: server.Router,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

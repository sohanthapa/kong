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
	server := &server.Server{
		Router: chi.NewRouter(),
	}
	//call to initialize server
	server.Initialize()

	//call to initialize mongodb connnection client
	client := mongostore.New()

	//initialize the data in db
	mongostore.Initialize(client)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", defaultPort),
		Handler: server.Router,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

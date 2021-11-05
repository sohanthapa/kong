package server

import (
	"kong/stores"
)

// Config is the configuration for the server
type Config struct {
	DataStore stores.Storer
}

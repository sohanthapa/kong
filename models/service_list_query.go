package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ServiceListQuery struct {
	Limit  int64
	Offset int64
	Sort   string
	Filter bson.M
}

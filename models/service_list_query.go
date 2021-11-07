package models

// ServiceListQuery represent a collection of parameters to query for services
type ServiceListQuery struct {
	Limit  int64
	Offset int64
	Sort   string
	Filter string
}

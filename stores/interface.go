package stores

import "kong/models"

// Storer is the aggregate interface that represents entity across the application
type Storer interface {
	ServiceStorer
}

type ServiceStorer interface {
	CreateService(*models.Service)
}

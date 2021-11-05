package data

import "kong/models"

var (
	Service1 = models.Service{
		ID:          "ServiceID1",
		Name:        "Service1",
		Description: "This is Service one",
	}
	Service2 = models.Service{
		ID:          "ServiceID2",
		Name:        "Service2",
		Description: "This is Service two",
	}
	Service3 = models.Service{
		ID:          "ServiceID3",
		Name:        "Service3",
		Description: "This is Service three",
	}
)

var ServiceList = []models.Service{Service1, Service2, Service3}

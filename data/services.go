package data

import "kong/models"

var (
	Service1 = models.Service{
		ID:          "ServiceID1",
		Name:        "Service1",
		Description: "This is Service one",
		Versions:    models.Versions{Version1, Version2},
	}
	Service2 = models.Service{
		ID:          "ServiceID2",
		Name:        "Service2",
		Description: "This is Service two",
		Versions:    models.Versions{Version3, Version4},
	}
	Service3 = models.Service{
		ID:          "ServiceID3",
		Name:        "Service3",
		Description: "This is Service three",
		Versions:    models.Versions{Version1, Version2},
	}
	Service4 = models.Service{
		ID:          "ServiceID4",
		Name:        "Service3",
		Description: "This is Service four",
	}
	Service5 = models.Service{
		ID:          "ServiceID5",
		Name:        "Service3",
		Description: "This is Service five",
	}
	Service6 = models.Service{
		ID:          "ServiceID6",
		Name:        "Service4",
		Description: "This is Service six",
		Versions:    models.Versions{Version2, Version4},
	}
)

var ServiceList = []models.Service{Service1, Service2, Service3, Service4, Service5, Service6}

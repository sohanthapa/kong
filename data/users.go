package data

import (
	"kong/models"
)

var (
	User1 = models.User{
		ID:          "User1",
		Name:        "TestUser1",
		Permissions: []string{models.PermGetService, models.PermGetServiceList},
	}

	User2 = models.User{
		ID:          "User2",
		Name:        "TestUser2",
		Permissions: []string{models.PermGetServiceList},
	}
)

// UserList stores the user slice
var UserList = []models.User{User1, User2}

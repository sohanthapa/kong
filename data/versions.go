package data

import "kong/models"

var (
	Version1 = models.Version{
		ID:   "VersionID1",
		Name: "Version1",
	}
	Version2 = models.Version{
		ID:   "VersionID2",
		Name: "Version2",
	}
	Version3 = models.Version{
		ID:   "VersionID3",
		Name: "Version3",
	}
	Version4 = models.Version{
		ID:   "VersionID4",
		Name: "Version4",
	}
)

var VersionList = []models.Version{Version1, Version2, Version3, Version4}

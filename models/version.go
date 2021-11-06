package models

//Version contains the info required for a Version
type Version struct {
	ID   string //unique value
	Name string
}

//Versions stores Version type slice
type Versions []Version

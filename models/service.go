package models

type Service struct {
	ID          string // unique value
	Name        string
	Description string
	Versions    []Version
}

type Services []Service

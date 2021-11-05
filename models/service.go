package models

type Service struct {
	Name        string
	Description string
	Versions    []Version
}

type Services []Service

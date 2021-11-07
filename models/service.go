package models

//Service contains the info required for a Service
type Service struct {
	ID          string // unique value
	Name        string
	Description string
	Versions    []Version `json:",omitempty"` //stores all the versions for this service
}

//Services stores Service type slice
type Services []Service

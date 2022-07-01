package model

type Tests struct {
	Base            // Explicitly specify the type to be uuid
	Name     string `json:"name" xml:"name" form:"name"`
	LastName string `json:"last_name" xml:"last_name" form:"last_name"`
}

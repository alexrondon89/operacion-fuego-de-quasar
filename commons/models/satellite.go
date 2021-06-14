package models

type Satellite struct {
	Name 		string		`json:"name"`
	Distance 	float32		`json:"distance"`
	Message 	[]string	`json:"message"`
}
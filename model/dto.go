package model

type GlobalResponse struct {
	Message string 	`json:"message"`
	Data 	any		`json:"data"`
	Errors	any		`json:"errors"`
}
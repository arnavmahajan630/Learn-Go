package models

type MyEvent struct {
	Name string `json:"name"`
	Age int 	`json:"age"`
}

type Myresponse struct {
	Message string 	`json:"message"`
}
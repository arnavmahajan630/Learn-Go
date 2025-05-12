package main

import (
	"fmt"
	"net/http"

	"github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/REST-API-With-MongoDB/controllers"
	"github.com/globalsign/mgo"
	"github.com/julienschmidt/httprouter"
)


func main() {
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)
	fmt.Println("Server started on Port 12345")
	 err := http.ListenAndServe("localhost:12345", router) // golang server runs at 8080
	 if(err != nil) {
		panic(err)
	 }
}

func getSession() *mgo.Session {
	// establishing connection with MongoDB	
	res, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err !=nil {
		fmt.Println(err)
	}
	return res
}

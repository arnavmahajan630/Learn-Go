package main

import (
	"net/http"

	"github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/REST-API-With-MongoDB/controllers"
	"github.com/globalsign/mgo"
	"github.com/julienschmidt/httprouter"
)


func main() {
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.Get.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", router)
}

func getSession() *mgo.Session {
	// establishing connection with MongoDB	
	res, err := mgo.Dial("mongodb://localhost:27107")
	if err !=nil {
		panic(err)
	}
	return res
}

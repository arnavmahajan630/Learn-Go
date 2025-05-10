package routes

import ("github.com/gorilla/mux"
        "github.com/Learn-Go/Simple-Projects/Book-Management-System/pkg/controllers"
	)

var RegisterStoreRzoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.HomePage()).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook()).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks()).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetbyId()).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook()).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook()).Methods("DELETE")
}
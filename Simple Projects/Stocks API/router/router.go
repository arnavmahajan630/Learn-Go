package router

import (
	controllers "github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/Stocks-API/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/stock/{id}", controllers.GetStockID).Methods("GET")
	router.HandleFunc("/api/stock", controllers.GetAllStocks).Methods("GET")
	router.HandleFunc("/api/stock/{id}", controllers.UpdateStocks).Methods("PUT")
	router.HandleFunc("/api/stock/{id}", controllers.DeleteStcok).Methods("DELETE")
	router.HandleFunc("/api/stock/{id}", controllers.CreateStock).Methods("POST")
	return router

}
package router

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/stock/{id}", middleware.GetStockID).Methods("GET")
	router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStocks).Methods("PUT")
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStcok).Methods("DELETE")
	router.HandleFunc("/api/stock/{id}", middleware.CreateStock).Methods("POST")
	return router

}
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "github.com/OceanWhisperer/pkg/models"
	"github.com/gorilla/mux"
)

var NewBook model.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newbooks := model.GetAllBooks()
	res, _ := json.Marshal(newbooks)
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) // a json list of books from db
}

func GetBookByID(w http.ResponseWriter, r * http.Request) {
   vars := mux.Vars(r)
   bid := vars["bookID"]
   id, err := strconv.ParseInt(bid,0,0)
  if err != nil {
	fmt.Println("Error Wile Parsing")
  }else {
	bookdeets, _ := model.GetBookByID(id)
	res , _ := json.Marshal(bookdeets)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
  }

}
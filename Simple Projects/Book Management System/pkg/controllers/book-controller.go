package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OceanWhisperer/pkg/models"
	"github.com/OceanWhisperer/pkg/utils"
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

func GetbyId(w http.ResponseWriter, r * http.Request) {
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

func CreateBook(w http.ResponseWriter, r * http.Request) {
	CreateBook := &model.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r * http.Request) {
	 vars := mux.Vars(r)
	 bid := vars["bookID"]
	 id, err:= strconv.ParseInt(bid,0,0)
	 if(err != nil) {
		fmt.Println("Error parsing ")
	 }
	  book := model.DeleteBook(id)
	  res, _ := json.Marshal(book)
	  w.Header().Set("Content-Type", "application/json")
	  w.WriteHeader(http.StatusOK)
	  w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r* http.Request) {
		var UpdateBook  = &model.Book{}
		utils.ParseBody(r, UpdateBook)
		vars := mux.Vars(r)
		bid := vars["bookID"]
		id, err := strconv.ParseInt(bid, 0,0)
		if( err == nil) {
			fmt.Println("error while Parsing")
		}
		bookdeets, _ := model.GetBookByID(id)
		if(UpdateBook.Name != "") {
			bookdeets.Name = UpdateBook.Name
		}
		if(UpdateBook.Author != "") {
			bookdeets.Author = UpdateBook.Author
		}
		if(UpdateBook.Publication != "") {
			bookdeets.Publication = UpdateBook.Publication	
		}
		res ,_ := json.Marshal(bookdeets)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
}
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/Stocks-API/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to Postgres Db")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode the body: ", err)
	}
	insertID := insertStock(stock)
	res := models.Resp{
		ID:      insertID,
		Message: "stock Created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks , err := getallstocks()

	if(err != nil){log.Fatal(err)}
	json.NewEncoder(w).Encode(stocks)

}

func GetStockID(w http.ResponseWriter, r* http.Request) {
    params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if(err != nil){log.Fatal(err)}
	stock := getstock(int64(id))
	json.NewEncoder(w).Encode(stock)


}

func UpdateStocks(w http.ResponseWriter, r* http.Request) {
  params := mux.Vars(r)
  id, err := strconv.Atoi(params["id"])
  if( err != nil) {log.Fatal(err)}
  var stock models.Stock
  err = json.NewDecoder(r.Body).Decode(&stock)
  if(err != nil) {
	log.Fatal(err)
  }
  updatedres := updatestocks(int64(id), stock)
  msg := fmt.Sprintf("Stock has been updated successfully : %v", updatedres)
  res := models.Resp{ID: int64(id), Message: msg, }
  json.NewEncoder(w).Encode(res)
}
func DeleteStcok(w http.ResponseWriter, r * http.Request) {
   params := mux.Vars(r)
   id, err := strconv.Atoi(params["id"])
   if( err != nil) {log.Fatal(err)}
   deletedres := deletestock(int64(id))
   fmt.Sprintf("The stock has been deleted successfully %v", deletedres)
   json.NewEncoder(w).Encode(deletedres)
}

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
	insertID := insertstock(stock)
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
	stock, _ := getstockid(int64(id))
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




func insertstock(stock models.Stock) int64{
  db := CreateConnection()
  defer db.Close()
  sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
  var id int64
  err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock,stock.Company).Scan(&id)
   if(err != nil){log.Fatal(err)}
   fmt.Printf("Inserted a single redcord %v", id)
   return id
}

func deletestock(id int64) int64{

}

func getstockid(id int64) (models.Stock, error){
  db := CreateConnection()
  defer db.Close()
  var stock models.Stock
  sqlstatement := `SELECT * FROM stocks WHERE stockid=$1`
  row := db.QueryRow(sqlstatement, id)
  err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

  switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return stock, nil
	case nil:
		return stock , nil
	
	default:
		log.Fatal("Unable to scan row %v", err)
  }
  return stock ,err
}

func getallstocks() ([]models.Stock, error){
   db := CreateConnection()
   defer db.Close()
   var stocks []models.Stock
   sqlstatement := `SELECT * FROM stocks`
   row, err  := db.Query(sqlstatement)
   if(err != nil) {
		log.Fatal("Unable to execute the error")
   }
   defer row.Close()
   for row.Next() {
	var stock models.Stock
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	if( err != nil){log.Fatal(err)}
	stocks = append(stocks, stock)
   }
   return stocks, err
}

func updatestocks(id int64, stock models.Stock) int64{

}
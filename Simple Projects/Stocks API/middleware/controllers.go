package middldeware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/Stocks-API/models"
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
   

}

func GetAllStocks() {

}

func GetStockID() {

}

func UpdateStocks() {

}
func DeleteStcok() {

}

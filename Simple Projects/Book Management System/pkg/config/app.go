package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Db Creation
var DB *gorm.DB

// Connect Function
func Connect() {
	dsn := "root:toor@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(" Failed to connect to MySQL: %v", err)
		panic(err)
	}

	fmt.Println(" Connected to MySQL database successfully.")
}

func GetDB() *gorm.DB {
	return DB
}

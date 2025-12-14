package database

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func DbInstance() * mongo.Client {
	err :=godotenv.Load(".env")
	if err != nil {
		
	}
}
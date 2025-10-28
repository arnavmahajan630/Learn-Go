package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/arnavmahajan630/Learn-Go/Intermediate-Projects/MyOwnDb/api"
	"github.com/arnavmahajan630/Learn-Go/Intermediate-Projects/MyOwnDb/db"
	"github.com/go-chi/chi/v5"
)

type User struct {
	Name string
	Age json.Number
	Contact string
	Company string
	Addy Address
}

type Address struct {
	City string
	State string
	Country string
	Pincode json.Number
}

func main() {
	fmt.Println("hello and testing")
	r := chi.NewRouter()
	r.Get("/", api.Welcome)

	directory := "./"
	if err := db.NewDb(directory, nil); err != nil {
		log.Fatal("Failed to create the database. Reason: %v", err)
	}

	emplace := []User{
		{"Jhon", "23", "1234567890", "NewComp", Address{"New York", "US", "US", "1221122112"}},
		{"Jhon", "22", "1234567890", "NewComp", Address{"Banglore", "Karnataka", "India", "1221122112"}},
		{"Doe", "28", "1234567890", "NewComp", Address{"New York", "US", "US", "1221122112"}},
		{"Joey", "23", "1234567890", "NewComp", Address{"New York", "US", "US", "1221122112"}},
		{"Will", "25", "1234567890", "NewComp", Address{"New York", "US", "US", "1221122112"}},
		{"Smith", "33", "1234567890", "NewComp", Address{"New York", "US", "US", "1221122112"}},
	}
	 // mock to be replaced with api
	for _, val := range emplace {
		db.Write("users", val.Name, User{
			Name: val.Name,
			Age: val.Age,
			Contact: val.Contact,
			Company: val.Company,
			Addy: val.Addy,
		})
	}

	records , err := db.ReadAll("users")
	if err !=  nil {
		log.Fatal("Failed to read from the Database. Reason: %v", err)
	}

	fmt.Println("records")

	
}
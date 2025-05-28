package main

import (
	db "github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/Simple-CRM/Database"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)  

func setupRoutes(app *fiber.App) {
    app.Get(getleads)
	app.Get(getlead)
	app.Post(postlead)
	app.Delete(deletelead)
}
func initDb() {
   var err error
   db.DbConn, err = gorm.Open("sqlite3", "simple_crm.db")
   
}
func main() {
	app := fiber.New() 
	initDb()
	setupRoutes(app)
	app.Listen(":3000")
	defer db.DbConn.Close()
}

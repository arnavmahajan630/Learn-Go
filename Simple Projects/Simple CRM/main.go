package main

import (
	"github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/Simple-CRM/database/database"
	"github.com/gofiber/fiber/v2"
)


func setupRoutes(app *fiber.App) {
	app.Get(Getleads)
	app.Get(getlead)
	app.Post(Newlead)
	app.Delete(Deletelead)
}

func initDB() {

}

func main() {
	app := fiber.New();
    app.Listen(":3000");

}
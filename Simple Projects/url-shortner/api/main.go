package main

import (
	"fmt"
	"log"

	"github.com/arnavmahajan630/Learn-Go/Simple-Proejcts/url-shortner/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"google.golang.org/genproto/googleapis/maps/routes/v1"
)

func setupRoutes(a * fiber.App) {
	a.Get("/:url", routes.ResolveUrl)
	a.Post("/api/v1", routes.ShortenUrl)
}


func main() {
	fmt.Println("hello world")
	cfg , err := config.Load()
	if err != nil {
		log.Fatal("Config error: ", err)
	}
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	log.Printf("Server Started on Port: %v", cfg.APP_PORT)
	log.Fatal(app.Listen(cfg.APP_PORT))
}
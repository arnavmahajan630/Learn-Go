package main

import (
	"net/http"

	"github.com/arnavmahajan630/Learn-Go/Simple-Projects/Golang-echo-api/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/health-check", handlers.HelthCheckerHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)
	e.GET("/", func(e echo.Context) error {return e.String(http.StatusOK, "Welcome to the server")})
	e.Logger.Fatal(e.Start(":3000"));
}
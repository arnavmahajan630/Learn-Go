package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelthCheckerHandler(e echo.Context) error{
	return e.String(http.StatusOK, "OK!"); 
}
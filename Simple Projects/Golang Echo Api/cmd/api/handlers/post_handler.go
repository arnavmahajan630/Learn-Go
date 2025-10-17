package handlers

import (
	"net/http"
	"strconv"

	util "github.com/arnavmahajan630/Learn-Go/Simple-Projects/Golang-echo-api/cmd/api/service"
	"github.com/labstack/echo/v4"
)

func PostIndexHandler(e echo.Context) error{
	data ,err := util.GetAll()
	if err != nil{
		e.String(http.StatusBadGateway, "Unable to process Data");
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return e.JSON(http.StatusOK, res);
}

func PostSingleHandler(e echo.Context) error {
	id := e.Param("id");
	idx , err := strconv.Atoi(id);
	if err != nil {
		return e.String(http.StatusBadGateway, "Unable to process")
	}
	data, err := util.GetByIdx(idx);
	if err != nil {
		return e.String(http.StatusBadGateway, "Unable to process")
	}
	res := make(map[string]any)
	res["status"] = "ok";
	res["data"] = data
	return e.JSON(http.StatusOK, data);
}
package main

import (
	"slices"

	"github.com/arnavmahajan630/Learn-Go/Intermediate-Projects/Pizza-Tracker/internal/models"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func RegistorCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_pizza_size", createSliceValidator(models.PizzaSizes))
		v.RegisterValidation("valid_pizza_type", createSliceValidator(models.PizzaTyles))
	}
}


func createSliceValidator(allwoedValues []string) validator.Func{
	return func(fl validator.FieldLevel) bool {
		return slices.Contains(allwoedValues, fl.Field().String())
	}
}
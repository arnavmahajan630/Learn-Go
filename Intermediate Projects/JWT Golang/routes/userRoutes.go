package routes

import (
	"github.com/arnavmahajan630/Learn-Go/Intermediate-Projects/JWT-Golang/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5/middleware"
)

func userRoutes(ir * gin.Engine) {
	ir.Use(middleware.Authenticate())
	ir.GET("/users", controllers.getUsers)
	ir.GET("/user/:userid", controllers.getUser)
}
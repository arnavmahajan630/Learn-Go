package routes

import (
	"github.com/arnavmahajan630/Learn-Go/Intermediate-Projects/JWT-Golang/controllers"
	"github.com/gin-gonic/gin"
)

func authRoutes(ir * gin.Engine) {
	ir.POST("/user/signup", controllers.Signup)
	ir.POST("/users/login", controllers.Login)

}
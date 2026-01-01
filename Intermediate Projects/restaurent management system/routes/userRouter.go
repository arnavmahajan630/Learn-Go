package routes

import (
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func ProtectedUserRoutes(inCommingRoutes *gin.Engine) {
	inCommingRoutes.GET("/users", controllers.GetUsers())
	inCommingRoutes.GET("/users/:user_id", controllers.GetUser())
}

func PublicUserRoutes(inCommingRoutes * gin.RouterGroup) {
	inCommingRoutes.POST("/users/singup", controllers.SignUp())
	inCommingRoutes.POST("/users/signin", controllers.SignIn())
}
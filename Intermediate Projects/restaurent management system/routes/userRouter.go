package routes

import (
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func ProtectedUserRoutes(inCommingRoutes * gin.RouterGroup) {
	inCommingRoutes.GET("/users", controllers.GetUsers())
	inCommingRoutes.GET("/users/:user_id", controllers.GetUser())
}

func PublicUserRoutes(inCommingRoutes *gin.Engine) {
	inCommingRoutes.POST("/users/singup", controllers.SignUp())
	inCommingRoutes.POST("/users/signin", controllers.SignIn())
}
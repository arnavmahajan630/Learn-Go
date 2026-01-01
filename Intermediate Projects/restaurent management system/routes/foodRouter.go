package routes

import (
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/controllers"
	"github.com/gin-gonic/gin"
)


func FoodRoutes(inCommingRoutes * gin.RouterGroup) {
	inCommingRoutes.GET("/foods", controllers.GetFoods())
	inCommingRoutes.GET("/food/:food_id", controllers.GetFood())
	inCommingRoutes.POST("/food", controllers.CreateFood())
	inCommingRoutes.PATCH("/foods/:food_id",controllers.UpdateFood())
}
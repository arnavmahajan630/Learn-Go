package routes

import (
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/controllers"
	"github.com/gin-gonic/gin"
)


func TableRoutes(inCommingRoutes * gin.RouterGroup) {
	inCommingRoutes.GET("/tables", controllers.GetTables())
	inCommingRoutes.GET("/tables/:order_id", controllers.GetTable())
	inCommingRoutes.POST("/tables", controllers.CreateTable())
	inCommingRoutes.PATCH("/tables/:table_id",controllers.UpdateTable())
}
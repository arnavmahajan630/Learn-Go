package routes

import "github.com/gin-gonic/gin"

func OrderRoutes(inCommingRoutes * gin.RouterGroup) {
	inCommingRoutes.GET("/orders", controllers.GetOrders())
	inCommingRoutes.GET("/orders/:order_id", controllers.GetOrder())
	inCommingRoutes.POST("/orders", controllers.CreateOrder())
	inCommingRoutes.PATCH("/orders/:order_id",controllers.UpdateOrder())
}
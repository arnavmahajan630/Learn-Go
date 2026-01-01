package routes

import (
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/controllers"
	"github.com/gin-gonic/gin"
)


func OrderItemRoutes(inCommingRoutes * gin.RouterGroup) {
	inCommingRoutes.GET("/orderItems", controllers.GetOrderItems())
	inCommingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	inCommingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	inCommingRoutes.POST("/orderItems", controllers.CreateMenu())
	inCommingRoutes.PATCH("/orderItems/:orderItem_id",controllers.UpdateMenu())

}
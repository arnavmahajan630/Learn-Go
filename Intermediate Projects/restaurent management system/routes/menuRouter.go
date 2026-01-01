package routes

import (
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(inCommingRoutes * gin.RouterGroup) {
	inCommingRoutes.GET("/menus", controllers.GetMenus())
	inCommingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	inCommingRoutes.POST("/menus", controllers.CreateMenu())
	inCommingRoutes.PATCH("/menus/:menus_id",controllers.UpdateMenu())
}
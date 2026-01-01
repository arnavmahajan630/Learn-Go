package main

import (
	"fmt"
	"os"

	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/database"
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/middleware"
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Public Registration and Login Routes
	routes.PublicUserRoutes(router)

	
	protected := router.Group("/")
	protected.Use(middleware.Authentication())

	// Protected Routes
	routes.ProtectedUserRoutes(protected)
	routes.FoodRoutes(protected)
	routes.MenuRoutes(protected)
	routes.InvoiceRoutes(protected)
	routes.TableRoutes(protected)
	routes.OrderRoutes(protected)
	routes.OrderItemRoutes(protected)

	// starting the server
	fmt.Println("Server started running on PORT: " + PORT)
	router.Run(":" + PORT)

}

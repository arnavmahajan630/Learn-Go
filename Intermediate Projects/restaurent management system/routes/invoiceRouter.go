package routes

import (
	"github.com/arnavmahajan630/Learn-Go/intermediate-projects/restaurent-management-system/controllers"
	"github.com/gin-gonic/gin"
)


func InvoiceRoutes(inCommingRoutes * gin.RouterGroup) {
	inCommingRoutes.GET("/invoices", controllers.GetInvoices())
	inCommingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	inCommingRoutes.POST("/invoices", controllers.CreateInvoice())
	inCommingRoutes.PATCH("/foods/:invoice_id",controllers.UpdateInvoice())
}
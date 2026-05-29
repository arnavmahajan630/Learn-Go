package main

import (
	"log/slog"
	"net/http"

	"github.com/arnavmahajan630/Learn-Go/Intermediate-Projects/Pizza-Tracker/internal/models"
	"github.com/gin-gonic/gin"
)

type OrderFormData struct {
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct {
	Name         string   `form:"name" binding:"required, min=2, max=100"`
	Phone        string   `form:"phone" binding:"required, min=10, max=10"`
	Address      string   `form:"address" binding:"required, min=10, max=100"`
	Sizes        []string `form:"size" binding:"required, min=1, dive,validate_pizza_size"`
	Types        []string `form:"type" binding:"required, min=1, dive,validate_pizza_type"`
	Instructions []string `form:"instructions" binding:"required, min=1"`
}


func (h * Handler) ServeNewOrderPost(c * gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", OrderFormData{
		PizzaTypes: models.PizzaTyles,
		PizzaSizes: models.PizzaSizes,
	})
}

func (h * Handler) SubmitOrderPost(c * gin.Context) {
	var form OrderRequest

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	orderItems := make([]models.OrderItem, len(form.Sizes))
	for i := range orderItems {
		orderItems[i] = models.OrderItem{
			Size: form.Sizes[i],
			Pizza : form.Types[i],
			Instructions: form.Instructions[i],

		}
	}

	order := models.Order {
		CustomerName: form.Name,
		Phone: form.Phone,
		Address: form.Address,
		Status: models.OrderStatus[0],
		Items: orderItems,

	}

	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("Failed to Create order", "error", err)
		c.String(http.StatusInternalServerError, "Something went Wrong")
	}
	slog.Info("Order Created", "orderID", order.ID, "customer", order.CustomerName)
	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)

}

func (h * Handler) ServeCustomer(c * gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		c.String(http.StatusBadRequest,  "Order Id is required")
	}

	order , err := h.orders.GetOrder(orderId)
	if err != nil {
		c.String(http.StatusNotFound,  "Order not found")
		return
	}

	c.HTML(http.StatusOK, "customer.tmpl", gin.H{
		"Order": order,

	})
	

}
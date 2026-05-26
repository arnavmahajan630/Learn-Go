package models

import (
	"time"

	"github.com/arnavmahajan630/Learn-Go/Intermediate-Projects/Pizza-Tracker/utils"
	"gorm.io/gorm"
)

var (
	OrderStatus = []string{"Order Placed", "Preparing", "Baking", "Quality Check", "Ready"}
	PizzaTyles  = []string{"Margarita", "Peperoni", "BBQ", "Meat Lovers", "Chicken Overload", "Supreme"}
	PizzaSizes  = []string{"Small", "Medium", "Large", "ExtraLarge"}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID           string      `gorm:"primaryKey; size:14" json:"id"`
	Status       string      `gorm:"not null" json:"status"`
	CustomerName string      `gorm:"not null" json:"customerName"`
	Phone        string      `gorm:"not null" json:"phone"`
	Address      string      `gorm:"not null" json:"address"`
	Items        []OrderItem `gorm:"foreignKey:OrderID" json:"pizzas"`
	CreatedAt    time.Time   `json:"createdAt"`
}

type OrderItem struct {
	ID           string `gorm:"primaryKey; size:14" json:"id"`
	OrderId      string `gorm:"index;" json:"orderId"`
	Size         string `gorm:"not null" json:"size"`
	Pizza        string `gorm:"not null" json:"pizza"`
	Instructions string `json:"instructions"`
}

func (o * Order) BeforeCreate(tx * gorm.DB) error{
	if(o.ID== "") {
		val, err := utils.GenerateShortID();
		if err != nil {
			return err
		}
		o.ID = val

	} 
	return nil
}

func (oi * OrderItem) BeforeCreate(tx * gorm.DB) error{
	if(oi.ID== "") {
		val, err := utils.GenerateShortID();
		if err != nil {
			return err
		}
		oi.ID = val

	} 
	return nil
}


func (o * OrderModel) CreateOrder(order * Order) error{
	return o.DB.Create(order).Error
}


func (o * OrderModel) GetOrder(id string) (*Order, error){
	var order Order
	err := o.DB.Preload("Items").First(&order, "id = ?", id).Error
	return &order, err
}

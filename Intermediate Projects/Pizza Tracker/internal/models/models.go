package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type OrderStorer interface {
	GetOrder(id string) (*Order, error)
	CreateOrder(order * Order) error
}


type dbModel struct {
	*OrderModel 
}

func InitDb(dataSouceName string) (OrderStorer, error) {
	db, err := gorm.Open(sqlite.Open(dataSouceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to Migrate database: %v", err)
	}

	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("Failed to automigrate: %v", err)
	}
	
	dbModel := dbModel{
		OrderModel: &OrderModel{DB: db},
	}
	return &dbModel, nil	
}
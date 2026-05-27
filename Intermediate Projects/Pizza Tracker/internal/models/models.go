package models

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type OrderStorer interface {
    GetOrder(id string) (*Order, error)
    CreateOrder(order *Order) error
}

// DbModel now holds the DB connection directly
type DbModel struct {
    DB *gorm.DB
}

func InitDb(dataSouceName string) (OrderStorer, error) {
    db, err := gorm.Open(sqlite.Open(dataSouceName), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect database: %v", err)
    }

    err = db.AutoMigrate(&Order{}, &OrderItem{})
    if err != nil {
        return nil, fmt.Errorf("failed to automigrate: %v", err)
    }
    
    // Return DbModel directly
    return &DbModel{DB: db}, nil    
}
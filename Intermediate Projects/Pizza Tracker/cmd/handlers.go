package main

import "github.com/arnavmahajan630/Learn-Go/Intermediate-Projects/Pizza-Tracker/internal/models"

type Handler struct {
	orders  models.OrderStorer
}


func Newhandler(Storage models.OrderStorer) *Handler {
	return &Handler{
		orders: Storage,
	}
}
package controllers

import "github.com/globalsign/mgo"

type UserController struct {
	session * mgo.Session
}

func NewUserController(m* mgo.Session) *UserController{
  return &UserController{m}
}

func GetUser() {
	
}

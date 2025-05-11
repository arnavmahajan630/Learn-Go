package models

import "github.com/globalsign/mgo/bson"

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Age    int8          `json:"age" bson:"age"`
	Gender string        `json:"gender" bson:"gender"`
}
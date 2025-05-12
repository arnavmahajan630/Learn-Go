package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/REST-API-With-MongoDB/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	session * mgo.Session
}

func NewUserController(m* mgo.Session) *UserController{
  return &UserController{m}
}


func (uc UserController) GetUser(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
    id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	obid := bson.ObjectIdHex(id)
	u := models.User{} // empty struct
	err := uc.session.DB("mongo-tut").C("users").FindId(obid).One(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Conetent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
   u := models.User{}
   json.NewDecoder(r.Body).Decode(&u)
   u.Id = bson.NewObjectId()
   err := uc.session.DB("mongo-tut").C("users").Insert(u)
   if(err != nil) {
	w.WriteHeader(404)
	return
   }
   res ,err := json.Marshal(u)
   if(err != nil) {
	fmt.Println(err)
   }else {
	w.Header().Set("Conetent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
   }
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	obid := bson.ObjectIdHex(id)
	err := uc.session.DB("mongo-tut").C("users").RemoveId(obid)
    if(err != nil) {
		w.WriteHeader(404)
		return;
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user ", obid, "\n")

}

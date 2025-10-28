package api

import "net/http"

func Welcome(w http.ResponseWriter,  r * http.Request) {
	w.Write([]byte("Welcome to this simple DataBase. hit /Add to add entries"))
}

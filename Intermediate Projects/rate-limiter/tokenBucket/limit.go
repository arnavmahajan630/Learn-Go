package main

import (
	"net/http"

	"golang.org/x/time/rate"
)


func rateLimiter(next func(w http.ResponseWriter, r * http.Request))http.HandlerFunc {
	limiter := rate.NewLimiter(2, 4)
	return http.HandleFunc(func(w http.ResponseWriter, r * http.Request){
		if !limiter.Allow()
			message := Message {
				Satus 
			}
	})	
	next(w, r)
}
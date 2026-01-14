package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

/*
Note on the closure patter
Here we take in the next function that we want to execute
we first write the code we want to execute before the og function
return a new handlerfunc if cases pass/ fail
then execute the function required
then if any more code required process it also.
*/

func rateLimiter(next func(w http.ResponseWriter, r * http.Request)) http.HandlerFunc{
	limiter := rate.NewLimiter(2, 4) // 2 req/sec total 4 tokens in bucket
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
		if !limiter.Allow(){
			message := Message {
				Satus: "Request Failed",
				Body: "Endpoint it at it's max capacity. Try in few secs",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			if err := json.NewEncoder(w).Encode(&message); err != nil {
				return
			}


		} else {
			next(w, r)
		}
	})
}
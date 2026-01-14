package main

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)


func perClientRateLimit(next func(w http.ResponseWriter, r * http.Request)) http.HandlerFunc{

	type Client struct {
		limiter * rate.Limiter
		lastSeen time.Time
	}

	var(
		mu sync.Mutex
		clients = make(map[string]*Client)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		mu.Lock()
		if _ , found := clients[ip]; !found {
			clients[ip] = &Client{limiter: rate.NewLimiter(2, 4)}
		}

		clients[ip].lastSeen = time.Now()
		allowed := clients[ip].limiter.Allow()
		mu.Unlock()
		if !allowed{
			message := Message {
				Satus: "Request Failed",
				Body: "Endpoint it at it's max capacity. Try in few secs",
			}	
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			if err := json.NewEncoder(w).Encode(&message); err != nil {
				return
			}
		}
		next(w, r)
	})
}
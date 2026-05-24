package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan any, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				fmt.Println("Getting Uri: ", url)
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan any)
	defer close(done)
	urls := []string{"https://google.com", "https://badd"}

	for r := range checkStatus(done, urls...) {
		if(r.Error != nil) {
			fmt.Printf("error: %v\n", r.Error)
			continue
		}
		fmt.Printf("Response %v\n", r.Response)
	}
}

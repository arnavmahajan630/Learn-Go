package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan any)

	go func() {
		time.Sleep(5* time.Second)
		close(ch)
	}()

	fmt.Println("Blocking on Read")
	select {
	case <- ch:
		fmt.Printf("Unblocking %v later\n", time.Since(startTime))
	}
}
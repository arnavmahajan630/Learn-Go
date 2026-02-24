package main

import (
	"fmt"
	"time"
)

func main() {

	var data int
	go func() {
		data++
	}()
	time.Sleep(50 * time.Microsecond)
	if data != 0 {
		fmt.Printf("the value is %v\n", data)
	}
}

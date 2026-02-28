package main

import (
	"fmt"
	"sync"

)

// Joinnig using waitGroups
func main() {

	var wg sync.WaitGroup

	sayHello := func () {
		defer wg.Done()
		fmt.Println("hello world")
	}

	wg.Add(1)
	go sayHello()
	fmt.Println("call from main")
	wg.Wait()
}
package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var count int
// 	var lock sync.Mutex
// 	var wg sync.WaitGroup

// 	increment := func() {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		count++
// 		fmt.Printf("incrementing: %d\n", count)
// 	}

// 	decrement := func() {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		count--
// 		fmt.Printf("Decrementing: %d\n", count)
// 	}

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			increment()
// 		}()
// 	}

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			decrement()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("All increment decrement operations complete !")

// }

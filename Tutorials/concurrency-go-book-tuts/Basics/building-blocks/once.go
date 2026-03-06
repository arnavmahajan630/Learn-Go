package main

// import (
// 	"fmt"
// 	"sync"
// )


// func main() {
	
// 	var count int = 0

// 	increment := func ()  {
// 		count++
// 	}
// 	decrement := func ()  {
// 		count--
// 	}
// 	var once sync.Once
// 	var twice sync.Once
// 	var wg sync.WaitGroup
// 	wg.Add(20)
// 	for i := 0; i < 20; i++ {
// 		go func ()  {
// 			defer wg.Done()
// 		once.Do(increment)
// 		// once.Do(decrement) 
// 		twice.Do(decrement)
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("Count is: %d\n", count)

// }
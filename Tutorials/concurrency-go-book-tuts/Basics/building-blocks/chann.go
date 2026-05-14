	package main

	import (
		"fmt"
		"sync"
	)

	func consumer(ch <- chan int,  wg * sync.WaitGroup) {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}

	func producer(ch chan <- int, wg * sync.WaitGroup) {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}

	func main() {

	var wg sync.WaitGroup
		wg.Add(2)
		ch := make(chan int)
		go producer(ch, &wg)
		go consumer(ch, &wg)
		wg.Wait()
		fmt.Println("Exchanged data successfully")
	}
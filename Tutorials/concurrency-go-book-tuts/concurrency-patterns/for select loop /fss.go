package main

import (
	"fmt"
	"math/rand"
)


func main() {

	newRandStream := func(done <-chan any) (<-chan int, <-chan any) {
		randStream := make(chan int)
		terminated := make(chan any)

		go func() {
			defer fmt.Println("Closing newRandStream Instance")
			defer close(randStream)
			defer close(terminated)

			for {
				select {
				case randStream <- rand.Int():
				case <- done: return
				}
			}
		}()

		return randStream, terminated
	}

	done := make(chan any)
	randomStream, terminated := newRandStream(done)

	fmt.Println("Printing 3 Digits")
	
	for i :=1; i <=3; i++ {
		fmt.Printf("%d. %v\n", i, <-randomStream)
	}
	close(done)	
	<-terminated
	fmt.Println("Done")
}
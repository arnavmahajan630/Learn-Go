package main

import (
	"fmt"
	"sync"
	"time"
)

func tee(done <-chan any, in <-chan int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range in {
			fmt.Printf("\n[Tee] Processing value %d\n", val)

			o1 := out1
			o2 := out2

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					fmt.Println("[TEE] cancelled")
					return
				case o1 <- val:
					fmt.Printf("[Tee] Sent %d to A\n", val)
					o1 = nil
				case o2 <- val:
					fmt.Printf("[Tee] Sent %d to B\n", val)
					o2 = nil

				}
			}
			fmt.Printf("[Tee] Finished duplicating %d\n", val)
		}
	}()

	return out1, out2
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	in := make(chan int)
	done := make(chan any)
	a, b := tee(done, in)

	// a fast consumer
	go func() {
		defer wg.Done()
		for v := range a {
			fmt.Printf("A received %d at %v", v, time.Now())
		}
	}()

	// a slow consumer
	go func() {
		
		defer wg.Done()
		for v := range b {
			fmt.Printf("B received %d at %v", v, time.Now())
		}
	}()

	// producer

	go func() {
		defer close(in)

		for i := 1; i <= 5; i++ {
			in <- i
			fmt.Printf("\n[PRODUCER] Sent %d\n", i)
			
		}
	}()
	
	fmt.Println("[Main] Waiting for Consumers To Complete Consuming")
	wg.Wait()
	fmt.Println("Compelted executing successfully")
}

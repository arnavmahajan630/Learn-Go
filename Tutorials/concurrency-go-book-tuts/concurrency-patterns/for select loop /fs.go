package main

import (
	"fmt"
	"time"
)

func main() {

	doWork := func(done chan any, strings chan any) <-chan any {
		terminated := make(chan any)

		go func() {
			defer fmt.Println("doWork Exited")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					// doing work
					fmt.Println(s)
				case <- done: 
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan any)
	terminated := doWork(done , nil)

	go func() {
		time.Sleep(1*time.Second)
		fmt.Println("cancelling the doWork goroutine")
		close(done)
	}()

	<-terminated
	fmt.Println("Done")
}
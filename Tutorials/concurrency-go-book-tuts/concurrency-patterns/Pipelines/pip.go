package main

import "fmt"

func main() {

	generator := func(done <-chan any, integers ...int) <-chan int {
		intStream := make(chan int, len(integers))

		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan any, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int, len(intStream))

		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(done <-chan any, intStream <-chan int, additive int) <-chan int {
		addStream := make(chan int, len(intStream))

		go func() {
			defer close(addStream)
			for i := range intStream {
				select {
				case <-done: return
				case addStream <- i + additive:
				}
			}
		}()
		return addStream
	}

	done := make(chan any)
	defer close(done)

	intStream := generator(done , 1,2,3,4)
	pipleline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipleline {
		fmt.Println(v)
	}

}

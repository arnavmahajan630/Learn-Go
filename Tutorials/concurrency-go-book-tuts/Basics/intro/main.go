package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu * sync.Mutex
	value int
}


func (c Counter) Increment(){
	c.mu.Lock()
	c.value++;
	c.mu.Unlock()
} // guarding the state of a stuct example. we go with the primitive type

func main() {

	var data int
	go func() {
		data++
	}()
	time.Sleep(1 * time.Microsecond)
	if data != 0 {
		fmt.Printf("the value is %v\n", data)
	} else {fmt.Println("Value was unchanged")}
}

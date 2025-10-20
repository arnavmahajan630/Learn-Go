package primitives

import (
	"fmt"
)

// func somefunc(num string) {
// 	fmt.Println(num)
// }

func PrimitiveCon() {
	// go somefunc("12")
    // go somefunc("13")
	// go somefunc("21")

	// fmt.Println("hello")
	// time.Sleep(time.Second)

	mychan := make(chan string) // a communication link between different goroutines
	hischan := make(chan string) // another channel that gives string

	go func () {  // ananymous concurrent function sends data to channel.
		mychan<- "data"
		//time.Sleep(time.Second *10); // don't care since 
	}()

	go func(){
		hischan <- "Coca Cola"
	}()

	//msg  := <- mychan // we read this data from channel. Blocking

	// select allows recieveing messages from multiple channels
	// it's only going to select 1 of these channels at random
	// not at all like it's reading from all these

	// select {
	// case msgfrommychan := <- mychan:
	// 	fmt.Println(msgfrommychan)
	// case msgfromhischan := <- hischan:
	// 	fmt.Println(msgfromhischan)
	// }

	recieved := []string{}

	for i := 0; i < 2; i++ {
		select {
		case msg := <- mychan:
			recieved = append(recieved, msg)
		case msg := <- hischan:
			recieved = append(recieved, msg)
		}
	}

	fmt.Println(recieved)

	/*
	select blocks a parent goroutine until any of the goroutines 
	defined in it return something from a channel. Whichever returns 
	first is reported, and select is broken. Now the parent can 
	continue execution.

	Must use a loop but we then need to know how many messages we need
	*/
}
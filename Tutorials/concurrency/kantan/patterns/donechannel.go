package patterns

import (
	"fmt"
	"time"
)

func dowork(donechan <-chan bool) { // recieves a mesage for stopping via donechan
		for {
			select {
			case <- donechan:
				return;
			default:
				fmt.Println("Infinite loop example")
			}
		}
}

func Forselect2() {

	donechan := make(chan bool)
	go dowork(donechan) // this is read only passing
	time.Sleep(4*time.Second)
	close(donechan) // sends a false and unblocks the work
}
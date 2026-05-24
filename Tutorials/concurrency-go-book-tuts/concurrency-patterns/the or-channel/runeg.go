package main

import (
	"fmt"
	"time"
)

func sig(after time.Duration, name string) <-chan any {
	c := make(chan any)

	go func() {
		time.Sleep(after)
		fmt.Printf("[SIGNAL] %s closed\n", name)
		close(c)
	}()

	return c
}

var id int

func or(channels ...<-chan any) <-chan any {
	id++
	myID := id

	fmt.Printf("Creating OR-%d with %d channels\n", myID, len(channels))

	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan any)

	go func() {
		defer func() {
			fmt.Printf("Closing orDone-%d\n", myID)
			close(orDone)
		}()

		switch len(channels) {
		case 2:
			fmt.Printf("OR-%d waiting on 2 channels\n", myID)

			select {
			case <-channels[0]:
				fmt.Printf("OR-%d woke from channel 0\n", myID)
			case <-channels[1]:
				fmt.Printf("OR-%d woke from channel 1\n", myID)
			}

		default:
			fmt.Printf("OR-%d waiting on 3 channels + child OR\n", myID)

			select {
			case <-channels[0]:
				fmt.Printf("OR-%d woke from channel 0\n", myID)
			case <-channels[1]:
				fmt.Printf("OR-%d woke from channel 1\n", myID)
			case <-channels[2]:
				fmt.Printf("OR-%d woke from channel 2\n", myID)
			case <-or(append(channels[3:], orDone)...):
				fmt.Printf("OR-%d woke from CHILD done\n", myID)
			}
		}
	}()

	return orDone
}

func main() {
	start := time.Now()

	fmt.Println("Starting OR demo...")

	<-or(
		sig(6*time.Second, "a"),
		sig(7*time.Second, "b"),
		sig(8*time.Second, "c"),
		sig(9*time.Second, "d"),
		sig(10*time.Second, "e"),
		sig(6*time.Second, "f"),
		sig(7*time.Second, "g"),
	)

	fmt.Printf("\nMain woke after %v\n", time.Since(start))
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker produces results while periodically sending heartbeats.
//
// There are three important communication channels:
//
// 1. heartbeat - tells the outside world that the worker is alive
// 2. results   - carries actual useful work
// 3. done      - tells the worker to shut down

func worker(
	done <-chan struct{},
	pulseInterval time.Duration,
	workInterval time.Duration,
) (<-chan struct{}, <-chan int) {

	// Buffered heartbeat channel.
	//
	// We only need to preserve one heartbeat.
	// Heartbeats are signals, not data.
	heartbeat := make(chan struct{}, 1)

	// Results are unbuffered because every result matters.
	results := make(chan int)

	go func() {

		// Always close channels when the goroutine exits.
		defer close(heartbeat)
		defer close(results)

		pulseTicker := time.NewTicker(pulseInterval)
		defer pulseTicker.Stop()

		workTicker := time.NewTicker(workInterval)
		defer workTicker.Stop()

		// Non-blocking heartbeat.
		sendHeartbeat := func() {

			select {

			// If someone is listening or there is room
			// in the buffer, send the heartbeat.
			case heartbeat <- struct{}{}:

			// Otherwise discard the heartbeat.
			//
			// NEVER block the worker just because
			// nobody cares about the heartbeat.
			default:
			}
		}

		for {

			select {

			// Someone requested shutdown.
			case <-done:
				return

			// Time to send a heartbeat.
			case <-pulseTicker.C:
				sendHeartbeat()

			// New unit of work.
			case <-workTicker.C:

				value := rand.Intn(100)

				// IMPORTANT:
				// The result is critical, so unlike the heartbeat,
				// this send IS allowed to block.
				select {

				case <-done:
					return

				case results <- value:
					// result successfully delivered
				}
			}
		}
	}()

	return heartbeat, results
}

func main() {

	const (
		heartbeatInterval = 500 * time.Millisecond
		workInterval      = 2 * time.Second
		healthTimeout     = 1500 * time.Millisecond
	)

	done := make(chan struct{})

	time.AfterFunc(10*time.Second, func() {
		close(done)
	})


	heartbeat, results := worker(
		done,
		heartbeatInterval,
		workInterval,
	)

	fmt.Println("monitor started")

	for {

		select {

		case _, ok := <-heartbeat:

			if !ok {
				fmt.Println("heartbeat channel closed")
				return
			}

			fmt.Println("[heartbeat] worker is alive")

		

		case result, ok := <-results:

			if !ok {
				fmt.Println("results channel closed")
				return
			}

			fmt.Printf("[result] received value: %d\n", result)


		case <-time.After(healthTimeout):

			fmt.Println("[ERROR] worker is not healthy!")
			return
		}
	}
}
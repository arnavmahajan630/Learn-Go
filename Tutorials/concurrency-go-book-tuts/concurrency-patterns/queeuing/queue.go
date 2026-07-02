package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID int
}

func cashier(orders chan<- Order, totalOrders int) {
	for i := 1; i <= totalOrders; i++ {

		order := Order{ID: i}

		fmt.Printf("[Cashier] Customer placed Order #%d\n", order.ID)

		orders <- order

		fmt.Printf("[Cashier] Order #%d added to queue\n", order.ID)

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n[Cashier] No more customers.")
	close(orders)
}

func barista(id int, orders <-chan Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range orders {

		fmt.Printf("    [Barista %d] Started Order #%d\n", id, order.ID)

		time.Sleep(3 * time.Second)

		fmt.Printf("    [Barista %d] Finished Order #%d\n", id, order.ID)
	}

	fmt.Printf("    [Barista %d] Going home.\n", id)
}

func main() {

	orders := make(chan Order, 3)

	var wg sync.WaitGroup

	wg.Add(2)

	go barista(1, orders, &wg)
	go barista(2, orders, &wg)

	cashier(orders, 10)

	wg.Wait()

	fmt.Println("\nCoffee shop closed.")
}
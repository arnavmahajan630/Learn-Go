package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	cond  = sync.NewCond(&mu)
	tasks []string
	wg sync.WaitGroup
)

func worker(id int) {
	defer wg.Done()
	cond.L.Lock()

	for len(tasks) == 0 {
		fmt.Println("Worker", id, "waiting for task")
		cond.Wait()
	}

	task := tasks[0]
	tasks = tasks[1:]

	cond.L.Unlock()

	fmt.Println("Worker", id, "processing:", task)
}

func main() {

	go worker(1)

	wg.Add(1)
	cond.L.Lock()
	tasks = append(tasks, "send email")
	fmt.Println("Task added")
	cond.Signal()
	cond.L.Unlock()

	wg.Wait()

	fmt.Println("Task processed successfully")
}
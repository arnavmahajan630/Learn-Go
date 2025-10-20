package main

import (
	"fmt"
	"time"
)

func somefunc(num string) {
	fmt.Println(num)
}

func main() {
	go somefunc("12")
    go somefunc("13")
	go somefunc("21")

	fmt.Println("hello")
	time.Sleep(time.Second)
}
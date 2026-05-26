package main

import (
	"fmt"
	"runtime"
)

func main() {
	nums := runtime.NumCPU()
	fmt.Println(nums)
}

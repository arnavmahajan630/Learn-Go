package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// taking input from the user

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("welcome to pokiDesk")
	fmt.Printf(">")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("echoing: " + input)

}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	// taking input from the user

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("welcome to pokiDesk")
	for {
		fmt.Printf(">")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println("echoing: " + input + "\n")
		
	}

}

func cleanInput(str string)[]string {
	lowered := strings.ToLower(str) // lowersize all text
	words := strings.Fields(lowered)// break down string into words for messy input
	return words
}
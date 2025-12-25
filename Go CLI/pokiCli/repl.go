package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func startRepl() {
	// taking input from the user

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("welcome to pokiDesk\n\n")
	for {
		fmt.Printf(">")
		scanner.Scan()
		input := scanner.Text()

		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}
		fmt.Printf("echoing: %v \n", cleaned)

		cmd := cleaned[0]

		switch cmd {
		case "exit":
			fmt.Println("Exiting pokiCli Thanks for using!")
			time.Sleep(2 * time.Second)
			os.Exit(0)
		case "help":
			hlp := `Welcome to the PokiDesk Help menu Here are your available commands :`
			fmt.Println(hlp)
			fmt.Println("--help")
			fmt.Println("--exit")
		default:
			fmt.Println("Invalid command please try again")

		}
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)  // lowersize all text
	words := strings.Fields(lowered) // break down string into words for messy input
	return words
}

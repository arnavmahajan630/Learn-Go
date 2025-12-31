package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Home() {

	banner := `___________    .__            __   
\__    ___/___ |  |__ _____ _/  |_ 
  |    |_/ ___\|  |  \\__  \\   __\
  |    |\  \___|   Y  \/ __ \|  |  
  |____| \___  >___|  (____  /__|  
             \/     \/     \/      `


	fmt.Println(banner)
	fmt.Printf("Type Something to start chatting :). CTRL+C to exit\n\n")

	scanner := bufio.NewScanner(os.Stdin)


	for {
		fmt.Printf(">")
		scanner.Scan()
		input := scanner.Text()
		cleaned := CleanInput(input)
		if len(cleaned) == 0 {
			fmt.Println("Please Enter a valid chat message")
			continue
		}

		finalInput := strings.Join(cleaned, " ")
		resp, err := Chat(finalInput)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("AI > %v\n", resp)
	}



}
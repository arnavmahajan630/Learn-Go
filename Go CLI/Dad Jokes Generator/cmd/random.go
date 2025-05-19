/*
Copyright © 2025 NAME HERE <Ocean Whisperer>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getrandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getrandomJoke() {
	resb := getJokeData("https://icanhazdadjoke.com/")
	var joke Joke
	err := json.Unmarshal(resb, &joke)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}
	fmt.Println("Joke: ", joke.Joke)

}

func getJokeData(BaseApi string) []byte {
	req, err := http.NewRequest(http.MethodGet, BaseApi, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Go-CLI-Dad-Jokes-Generator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer res.Body.Close()
	resb, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	return resb
}

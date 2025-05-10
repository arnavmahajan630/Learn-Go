package main

import (
	"fmt"
	"io"
	"net/http"
)

func check(err error) {
	if(err != nil) {
		panic(err);
	}
}

func main() {
	fmt.Println("hello world")
	resp , err := http.Get("https://jsonplaceholder.typicode.com/todos/")
	check(err)
	body , err := io.ReadAll(resp.Body)
	check((err))
	fmt.Println("The data is: ", string(body))

}
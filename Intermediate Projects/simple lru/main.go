package main

import (
	"fmt"

	cache "github.com/arnavmahajan630/learn-go/intermidiate-projects/simple-lru/Cache"
)

func main() {
	fmt.Println("Simple LRU Cache Implementation")

	mycache := 	cache.NewCache()	

	for _, x := range []string{"parrot", "avocado", "dragonfruit", 
	"tree", "potato", "tomato"} {
		mycache.Check(x)
		mycache.Display()
	}
 }
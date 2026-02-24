package main

import (
	"fmt"
	"sort"
)


func add(i int , j int ) int {return i + j}
func sub(i int , j int ) int {return i - j}
func mult(i int , j int ) int {return i * j}
func div(i int , j int ) int {return i / j}

// var mp = map[string]func(int, int)int{
// 	"+":add,
// 	"-":sub,
// 	"*":mult,
// 	"/":div,
// }

type calFuncType func(int, int )int // calculator function type as mentioned above

var mp = map[string]calFuncType{
	"+":add,
	"-":sub,
	"*":mult,
	"/":div,
}

// Anonymous Functions

func main() {
	fmt.Println("hello world")

	for i := 0; i < 5; i++ {
		
		func(j int ){
			fmt.Println("printing", j, "from inside of anonymous function")
		}(i) // this serves as the argument i.e the value for j = i

	}


	// sorting example run

	type person struct {
		Fname string 
		Lname string
	}

	slice := []person{
		{"arnav", "mahajan"},
		{"james", "Doakes"},
		{"Wolf", "abhram"},
	}

	sort.Slice(slice, func(i, j int) bool {return slice[i].Lname < slice[j].Lname})
	// sort(v.begin(), v.end() [int i , int j]{return i < j. type shi});

	
}
package main

import (
	"fmt"

	"github.com/arnavmahajan630/learn-go/tutorials/abstraction/abstraction-kantan/vm"
)

type VendingMachine interface {
	GetDrink(money int64, brand string) string
}

type Application struct {
	vm VendingMachine
}

func (a Application) Run() {
	mydrink := a.vm.GetDrink(100, "cola")
	fmt.Println(mydrink)
}

func newApplication(vm VendingMachine) *Application {
	return &Application{vm: vm}
}

func main() {
	vendingmachine := vm.New()
	app := newApplication(vendingmachine)
	app.Run()
}
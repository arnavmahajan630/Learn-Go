package main

import (
	"fmt"
	"log"
	"practice-error-propogation/service"
)


func Handler() {
	data , err := service.DownloadInvoice("42")
	if err != nil {
		if _, ok := err.(service.ServiceErorr); ok {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Unexpected Error. Contact Support")

		log.Printf("%#v\n", err)

        return
	}

	fmt.Println(string(data))

}
func main() {
	Handler()
}

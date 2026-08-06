package main

import (
	"errors"
	"fmt"
	"log"
	something "practice-error-propogation/err"
	"practice-error-propogation/service"
)


func Handler() {
	data , err := service.DownloadInvoice("42")
	if err != nil {
		// Developer LOG: Complete error formatted
		if _, ok := err.(service.ServiceErorr); ok {
			fmt.Println(err.Error())
		}

		// User Facing
		var appErr something.AppError
		if errors.As(err, &appErr) {
			fmt.Println(appErr.SafeMessage())
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

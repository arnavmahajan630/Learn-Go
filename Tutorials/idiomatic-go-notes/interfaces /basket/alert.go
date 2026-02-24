package alert

import "fmt"

type EmailAltert struct {}
type SmsAlert struct {}

func(ea EmailAltert) Alert(message string) error{
	fmt.Println("Alerted message via email: ", message)
	return nil
}

func (sa SmsAlert)Alert(message string)error {
	fmt.Println("Alerted via SMS", message)
	return nil
}


package main

import (
	"fmt"
	alert "interfaces/basket"
)

type Alterter interface {
	Alert(message string)error
}

type WeatherStation struct {
	al Alterter 
}
func (w WeatherStation) ReportWeather(weather string) {
	s := fmt.Sprintln("WARNING: Extreme weather detected", weather)
	if err := w.al.Alert(s); err != nil {
		fmt.Printf("Failed to send the message")
	}
	fmt.Printf("Message sent successfully!")
}

func main() {
	//mailalert := alert.EmailAltert{}
	smsAlert := alert.SmsAlert{}
	ws := WeatherStation{smsAlert}
	ws.ReportWeather("Extreme Tornado")
}
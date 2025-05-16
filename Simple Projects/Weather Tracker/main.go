package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type ApiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temperature"`
	} `json:"main"`
}

func loadconfig(filename string) (ApiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return ApiConfigData{}, err
	}
	var c ApiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return ApiConfigData{}, err
	}
	return c, nil

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Go Weather Getter"))
}

func querry(city string) (WeatherData, error) {
	apconfig, err := loadconfig(".apiConfig")
	if err != nil {
		log.Fatal(err)
		return WeatherData{}, err
	}
	resp , err := http.Get("http://api.openweathermap.org/data/3.0/weather/APPID=" + apconfig.OpenWeatherMapApiKey + "&q" + city)
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()
	var d WeatherData
	err = json.NewDecoder(resp.Body).Decode(&d)
	if( err != nil) {
		return WeatherData{}, err
	}

	return d, nil

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		d, err := querry(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(d)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

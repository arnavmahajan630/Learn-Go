package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	PhotoApi = "https://api.pexels.com/v1/"
	VideoApi = "https://api.pexels.com/videos/"
)

type client struct {
	TOKEN          string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string) *client {
	c := http.Client{}
	return &client{TOKEN: token, hc: c}
}

type SearchResult struct {
	Page        int32   `json:"page"`
	PerPage     int32   `json:"perpage"`
	TotalResult int32   `json:"totalresults"`
	NextPage    string  `json:"nextpage"`
	Photos      []Photo `json:"photos"`
}

type Photo struct {
	Id   int32 `json:"id"`
	Width int32 `json:"width"`
	Height int32`json:"height"`
	URL string`json:"url"`
	Src PhotoSource  `json:"src"`
}

type PhotoSource struct {
	Orignal string `json:"orignal"`
	Large string `json:"large"`
	Large2x string `json:"large2x"`
	Medium string `json:"medium"`
	Small string `json:"small"`
	Potrait string `json:"potrait"`
	Square string `json:"square"`
	Landscape string `json:"landscape"`
	Tiny string `json:"tiny"`
}

func (c client) SearchPhotos(querry string, perpage int, page int) (*SearchResult, error){
  
}



func main() {
	os.Setenv("PexelsToken", "FRpq7clpC4p23iDLSGJ1I6kMw9XO0ItgKg6A6x2HQzbCdasYL70ddckL")
	var TOKEN = os.Getenv("PexelsToken")
	var client = NewClient(TOKEN)

	res, err := client.SearchPhotos("waves")
	if err != nil {
		fmt.Errorf("Search Error: %v", err)
	}
	if res.Page == 0 {
		fmt.Errorf("Serch result is Wrong")
	}
	fmt.Println(res)
}

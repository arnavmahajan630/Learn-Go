package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
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
	Id     int32       `json:"id"`
	Width  int32       `json:"width"`
	Height int32       `json:"height"`
	URL    string      `json:"url"`
	Src    PhotoSource `json:"src"`
}

type PhotoSource struct {
	Orignal   string `json:"orignal"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Potrait   string `json:"potrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type CuratedResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	NextPage string  `json:"next_page"`
	Photos   []Photo `json:"photos"`
}

func (c *client) requestWithAuth(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.TOKEN)
	resp, err := c.hc.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	times, err := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
	if err != nil {
		return resp, err
	}
	c.RemainingTimes = int32(times)
	return resp, nil

}

func (c *client) GetRandomPhoto() (*Photo , error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	randNum := r.Intn(1000)
	 res , err := c.CuratedPhotos(1, randNum)
	 if err == nil && len(res.Photos) == 1 {return &res.Photos[0], nil}
	 return nil, err

}
func (c *client) SearchPhotos(querry string, perpage int, page int) (*SearchResult, error) {
	url := fmt.Sprintf(PhotoApi+"search?query=%s&per_page=%d&page=%d", querry, perpage, page)
	resp, err := c.requestWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(data, &result)
	return &result, err

}

func (c *client) CuratedPhotos(per_page, page int) (*CuratedResult, error) {
	url := fmt.Sprintf(PhotoApi+"curated?per_page=%d&page=%d", per_page, page)
	resp, err := c.requestWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result CuratedResult
	err = json.Unmarshal(data, &result)
	return &result, err

}

func (c *client) GetPhoto(id int32) (*Photo, error) {
	url := fmt.Sprintf(PhotoApi+"photos/%d", id)
	resp, err := c.requestWithAuth("GET", url)
	if err != nil {return nil, err}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {return nil, err}
	var result Photo
	err = json.Unmarshal(data, &result)
	if err != nil {return nil, err}
	return &result, nil
}

func main() {
	os.Setenv("PexelsToken", "FRpq7clpC4p23iDLSGJ1I6kMw9XO0ItgKg6A6x2HQzbCdasYL70ddckL")
	var TOKEN = os.Getenv("PexelsToken")
	var client = NewClient(TOKEN)

	res, err := client.SearchPhotos("waves", 2, 2)
	if err != nil {
		fmt.Printf("Search Error: %v", err)
	}
	if res.Page == 0 {
		fmt.Printf("Serch result is Wrong")
	}
	fmt.Println(res)
}

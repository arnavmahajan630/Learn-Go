package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}

// file reading function
func raw() ([] Data, error) {
	r, err := os.ReadFile("model/data.json")
	if err != nil {
		return nil, fmt.Errorf("could not read the file")
	}
	var payload Payload
	err = json.Unmarshal(r, &payload.Data)
	if err != nil {
		return nil, err
	}
	return payload.Data , nil
}

func GetAll() ([]Data, error) {
	data , err := raw()
	if err != nil {
		return data, nil
	}
	return data, nil
	
}

func GetByIdx(idx int ) (any, error) {
	data, err := raw();
	if err!=nil {
		return nil, err
	}
	if idx > len(data) {
		res := make([]string, 0);
		return res, nil
	}
	return data[idx], nil
}

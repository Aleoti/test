package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Pole struct {
	Type string `json:"type"`
}
type Slov struct {
	S []map[string]interface{}
}

func main() {
	response, err := http.Get("https://api.github.com/users/Aleoti/events/public")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Pole
	err = json.Unmarshal([]byte(body), &data)
	//fmt.Println(data)
	for _, v := range data {
		fmt.Println(v.Type)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	UserID    int64  `json:"userId"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos/1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {

	}
	defer res.Body.Close()

	var response Response
	json.NewDecoder(res.Body).Decode(&response)

	fmt.Println(response)
}

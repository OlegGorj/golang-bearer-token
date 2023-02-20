package main

import (
	"fmt"
	"net/http"
)

func main() {
	token := "my_bearer_token_from_first_step"

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// main logic goes here..

}

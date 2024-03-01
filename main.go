package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("gobaseball")

	// api url
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "statsapi.mlb.com",
		Path:   "/api/v1",
	}

	// add the endpoint to the url
	endpoint := fmt.Sprintf("%s/%s", baseUrl.String(), "schedule/games/?sportId=1")

	// http client to make the request
	client := http.DefaultClient

	// make request
	res, err := client.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	// check that request is ok
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: http code, %v", res.StatusCode)
	}

	// read the body of the response
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// print the response body
	fmt.Println(string(body))

}

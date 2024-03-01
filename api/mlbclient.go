package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type MlbClient struct {
	client  *http.Client
	baseUrl *url.URL
}

// singleton
var mlbClient *MlbClient

func NewMlbClient() *MlbClient {
	if mlbClient == nil {
		mlbClient = &MlbClient{
			client: http.DefaultClient,
			baseUrl: &url.URL{
				Scheme: "https",
				Host:   "statsapi.mlb.com",
				Path:   "/api/v1",
			},
		}
	}
	return mlbClient
}

// func (mlb *MlbClient) Get(endpoint string, params map[string]string) []byte {
func (mlb *MlbClient) Get(endpoint string) []byte {
	// TODO figure out if I should return error from this
	// or just let the program die

	// mlb.baseUrl.RawQuery = params.Encode()
	url := mlb.baseUrl.Query()
	url.Add("sportId", "1")

	mlb.baseUrl.RawQuery = url.Encode()

	// make request
	res, err := mlb.client.Get(mlb.baseUrl.String())
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

	return body
}

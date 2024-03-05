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
				Path:   "/api/v1/",
			},
		}
	}
	return mlbClient
}

func (mlb *MlbClient) Get(endpoint string, params map[string]string) []byte {
	mlbUrl := mlb.baseUrl
	mlbUrl.Path += endpoint

	queryValues := url.Values{}
	for key, value := range params {
		queryValues.Add(key, value)
	}
	mlbUrl.RawQuery = queryValues.Encode()

	// make request
	fmt.Printf("making request to: %s\n", mlbUrl.String())
	res, err := mlb.client.Get(mlbUrl.String())
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

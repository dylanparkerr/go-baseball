package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dylanparkerr/go-baseball/api"
)

func main() {
	// data := getGames()
	data := getTeam("111")
	prettyData := prettyPrint(data)
	fmt.Println(string(prettyData))
}

func prettyPrint(input []byte) []byte {
	var output bytes.Buffer
	err := json.Indent(&output, input, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return output.Bytes()
}

func getGames() []byte {
	mlb := api.NewMlbClient()

	params := make(map[string]string)
	params["sportId"] = "1"

	games := mlb.Get("schedule/games/", params)

	return games
}

func getTeam(team string) []byte {
	mlb := api.NewMlbClient()

	endpoint := "teams/" + team

	params := make(map[string]string)
	params["sportId"] = "1"

	games := mlb.Get(endpoint, params)

	return games
}

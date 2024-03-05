package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dylanparkerr/go-baseball/api"
)

func main() {
	// data := getGames()
	// data := getTeam("144")
	// data := getMeta("statTypes")
	data := getPlayers()
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

func getPlayers() []byte {
	mlb := api.NewMlbClient()
	params := make(map[string]string)
	params["sportId"] = "1"
	return mlb.Get("sports/1/players/", params)
}

func getMeta(meta string) []byte {
	mlb := api.NewMlbClient()
	params := make(map[string]string)
	params["sportId"] = "1"
	return mlb.Get(meta, params)
}

func getGames() []byte {
	mlb := api.NewMlbClient()
	params := make(map[string]string)
	params["sportId"] = "1"
	return mlb.Get("schedule/games/", params)
}

func getTeam(team string) []byte {
	mlb := api.NewMlbClient()
	endpoint := "teams/" + team
	params := make(map[string]string)
	params["sportId"] = "1"
	return mlb.Get(endpoint, params)
}

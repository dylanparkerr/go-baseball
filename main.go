package main

import (
	"fmt"
	"github.com/dylanparkerr/go-baseball/api"
)

func main() {
	fmt.Println("lets gobaseball")
	mlb := api.NewMlbClient()
	b := mlb.Get("games")
	fmt.Println(b)
}

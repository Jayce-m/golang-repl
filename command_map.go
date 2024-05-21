package main

import (
	"fmt"
	"log"

	"github.com/Jayce-m/pokedexcli/internal/pokeapi"
)

func commandMap() {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

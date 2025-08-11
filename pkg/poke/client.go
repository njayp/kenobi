package poke

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func GetRandomPokemonFromAPI() Pokemon {
	return GetPokemonFromAPI(rand.Intn(150) + 1) // Random ID from 1 to 150
}

func GetPokemonFromAPI(id int) Pokemon {
	// Get 3 random Pokemon (from the first 150 - original Pokemon!)
	pokemonID := rand.Intn(150) + 1 // Random ID from 1 to 150
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemonID)

	// Make HTTP request
	println("Fetching Pokemon from API:", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse JSON
	var poke Pokemon
	err = json.Unmarshal(body, &poke)
	if err != nil {
		panic(err)
	}

	return poke
}

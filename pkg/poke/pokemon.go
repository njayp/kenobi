package poke

import "fmt"

// Pokemon represents the data we get from PokeAPI
type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (p Pokemon) Print() string {
	return fmt.Sprintf("Name: %s, Height: %d, Weight: %d, Types: %v, Stats: %v",
		p.Name,
		p.Height,
		p.Weight,
		p.Types,
		p.Stats,
	)
}

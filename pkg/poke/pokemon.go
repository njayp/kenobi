package poke

import "fmt"

// Pokemon represents the data we get from PokeAPI
type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []Stat `json:"stats"`
	Types  []Type `json:"types"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Type struct {
	Name string `json:"name"`
}

func (p Pokemon) Print() string {
	return fmt.Sprintf("Name: %s, Height: %d, Weight: %d, Types: %+v, Stats: %+v",
		p.Name,
		p.Height,
		p.Weight,
		p.Types,
		p.Stats,
	)
}

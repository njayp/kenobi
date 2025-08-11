package main

import (
	"github.com/njayp/hades/pkg/lesson/one"
	"github.com/njayp/hades/pkg/poke"
)

func main() {
	// printPoke()
	one.Run()
}

func printPoke() {
	poke.GetRandomPokemonFromAPI().Print()
}

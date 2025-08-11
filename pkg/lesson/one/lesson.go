package one

import (
	"fmt"

	"github.com/njayp/hades/pkg/poke"
)

func Run() {
	fmt.Println("🎮 Pokemon Trainer School - Lesson 1")
	fmt.Println("=====================================")
	fmt.Println()

	// PART 1: Working with a single Pokemon (Structs)
	fmt.Println("📚 PART 1: Meet Your Starter Pokemon")
	fmt.Println("-------------------------------------")

	// Get your starter Pokemon
	myStarter := poke.GetStarterPokemon()

	name := getPokemonName(myStarter)
	fmt.Printf("Your starter: %s\n", name)

	stat := getHighestStat(myStarter)
	fmt.Printf("Highest stat: %s with value %d\n", stat.Stat.Name, stat.BaseStat)

	sum := getStatsSum(myStarter)
	fmt.Printf("Sum of stats: %d\n", sum)

	fmt.Println()

	// PART 2: Working with multiple Pokemon (Arrays)
	fmt.Println("📚 PART 2: Your First Pokemon Team")
	fmt.Println("-----------------------------------")

	// Get your first 3 Pokemon
	myTeam := []poke.Pokemon{
		poke.GetRandomPokemonFromAPI(),
		poke.GetRandomPokemonFromAPI(),
		poke.GetRandomPokemonFromAPI(),
	}

	names := getTeamNames(myTeam)
	for i, name := range names {
		fmt.Printf("Team member %d: %s\n", i+1, name)
	}

	mon := getStrongestPoke(myTeam)
	fmt.Printf("💪 Your strongest Pokemon is: %s\n", mon.Name)

	fmt.Println()

	// PART 3: Pokemon Type Analysis
	fmt.Println("📚 PART 3: Understanding Pokemon Types")
	fmt.Println("--------------------------------------")

	if hasFireType(myTeam) {
		fmt.Println("🔥 You have a fire type Pokemon!")
	} else {
		fmt.Println("❌ No fire type Pokemon on your team")
	}

	count := countWaterType(myTeam)
	fmt.Printf("💧 You have %d water type Pokemon in your team\n", count)

	fmt.Println()
}

func getPokemonName(p poke.Pokemon) string {
	return "" // Replace this with your implementation
}

func getHighestStat(p poke.Pokemon) poke.Stat {
	return poke.Stat{} // Replace this with your logic
}

func getStatsSum(p poke.Pokemon) int {
	return 0 // Replace this with your calculation
}

func getTeamNames(team []poke.Pokemon) []string {
	return nil // Replace with your implementation
}

func getStrongestPoke(team []poke.Pokemon) poke.Pokemon {
	return poke.Pokemon{} // Replace with your implementation
}

func hasFireType(team []poke.Pokemon) bool {
	return false // Replace with your implementation
}

func countWaterType(team []poke.Pokemon) int {
	return 0 // Replace with your implementation
}

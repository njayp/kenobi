package one

import (
	"fmt"
)

// Pokemon represents a Pokemon with its basic stats
type Pokemon struct {
	Name   string
	Type   string // "fire", "water", "grass", "electric", etc.
	Level  int
	HP     int // Health Points
	Attack int // Attack stat
}

func Run() {
	fmt.Println("🎮 Pokemon Trainer School - Lesson 1")
	fmt.Println("=====================================")
	fmt.Println()

	// PART 1: Working with a single Pokemon (Structs)
	fmt.Println("📚 PART 1: Meet Your Starter Pokemon")
	fmt.Println("-------------------------------------")

	// Get your starter Pokemon
	myStarter := getStarterPokemon()

	// TODO: Task 1 - Print your starter's name
	printPokemonName(myStarter)

	// TODO: Task 2 - Check if your Pokemon is strong (level >= 10)
	if isPokemonStrong(myStarter) {
		fmt.Println("💪 Your Pokemon is strong!")
	} else {
		fmt.Println("🌱 Your Pokemon needs more training!")
	}

	// TODO: Task 3 - Calculate battle power
	power := calculateBattlePower(myStarter)
	fmt.Printf("⚔️  Battle Power: %d\n", power)

	fmt.Println()

	// PART 2: Working with multiple Pokemon (Arrays)
	fmt.Println("📚 PART 2: Your First Pokemon Team")
	fmt.Println("-----------------------------------")

	// Get your first 3 Pokemon
	myTeam := getFirstTeam()

	// TODO: Task 4 - Print all Pokemon names using a for loop
	printTeamNames(myTeam)

	// TODO: Task 5 - Find your highest level Pokemon
	highest := findHighestLevel(myTeam)
	fmt.Printf("⭐ Highest level Pokemon: %s (Level %d)\n", highest.Name, highest.Level)

	// TODO: Task 6 - Count how many Pokemon are ready for battle (HP > 20)
	readyCount := countBattleReady(myTeam)
	fmt.Printf("✅ Pokemon ready for battle: %d out of 3\n", readyCount)

	fmt.Println()

	// PART 3: Pokemon Type Analysis
	fmt.Println("📚 PART 3: Understanding Pokemon Types")
	fmt.Println("--------------------------------------")

	// TODO: Task 7 - Check if you have a fire type Pokemon
	if hasFireType(myTeam) {
		fmt.Println("🔥 You have a fire type Pokemon!")
	} else {
		fmt.Println("❌ No fire type Pokemon on your team")
	}

	// TODO: Task 8 - Find your water type Pokemon (if any)
	waterPokemon := findWaterType(myTeam)
	if waterPokemon.Name != "" {
		fmt.Printf("💧 Your water Pokemon: %s\n", waterPokemon.Name)
	} else {
		fmt.Println("💧 No water Pokemon found")
	}

	fmt.Println()
}

// ===== ALREADY IMPLEMENTED FUNCTIONS - DO NOT MODIFY =====

// getStarterPokemon returns your first Pokemon
func getStarterPokemon() Pokemon {
	return Pokemon{
		Name:   "Charmander",
		Type:   "fire",
		Level:  5,
		HP:     39,
		Attack: 52,
	}
}

// getFirstTeam returns your first team of 3 Pokemon
func getFirstTeam() [3]Pokemon {
	team := [3]Pokemon{
		{Name: "Charmander", Type: "fire", Level: 12, HP: 39, Attack: 52},
		{Name: "Squirtle", Type: "water", Level: 10, HP: 44, Attack: 48},
		{Name: "Pidgey", Type: "flying", Level: 8, HP: 40, Attack: 45},
	}
	return team
}

// getRandomPokemonFromAPI fetches 3 random Pokemon from PokeAPI
// THIS FUNCTION IS PROVIDED - DO NOT MODIFY

// ===== FUNCTIONS TO IMPLEMENT (PART 1-3) =====

// TODO TASK 1: Print the Pokemon's name
// This is your first function! Just print the Pokemon's name.
// Expected output: "Your starter: Charmander"
func printPokemonName(p Pokemon) {
	// YOUR CODE HERE
	// Hint: Use fmt.Printf or fmt.Println
	// Hint: Access the name with p.Name
}

// TODO TASK 2: Check if a Pokemon is strong
// A Pokemon is considered strong if its level is 10 or higher
// Return true if strong, false if not
func isPokemonStrong(p Pokemon) bool {
	// YOUR CODE HERE
	// Hint: Check if p.Level >= 10
	// Hint: Return true or false

	return false // Replace this with your logic
}

// TODO TASK 3: Calculate battle power
// Battle power = Level * Attack / 10
// This is a simple formula to estimate Pokemon strength
func calculateBattlePower(p Pokemon) int {
	// YOUR CODE HERE
	// Hint: Multiply p.Level by p.Attack
	// Hint: Divide the result by 10
	// Hint: Return the final number

	return 0 // Replace this with your calculation
}

// TODO TASK 4: Print all Pokemon names in your team
// Use a for loop to print each Pokemon's name
// Expected output:
// "Team member 1: Charmander"
// "Team member 2: Squirtle"
// "Team member 3: Pidgey"
func printTeamNames(team [3]Pokemon) {
	// YOUR CODE HERE
	// Hint: Use a for loop with range
	// Hint: for i, pokemon := range team { ... }
	// Hint: Print both the position (i+1) and pokemon.Name
}

// TODO TASK 5: Find the highest level Pokemon
// Loop through the team and find which Pokemon has the highest level
func findHighestLevel(team [3]Pokemon) Pokemon {
	// YOUR CODE HERE
	// Hint: Start with highest := team[0]
	// Hint: Loop through team[1] and team[2]
	// Hint: If you find a higher level, update highest
	// Hint: Return the highest level Pokemon

	return Pokemon{} // Replace with your implementation
}

// TODO TASK 6: Count how many Pokemon are ready for battle
// A Pokemon is ready if its HP is greater than 20
func countBattleReady(team [3]Pokemon) int {
	// YOUR CODE HERE
	// Hint: Start with count := 0
	// Hint: Loop through each Pokemon
	// Hint: If pokemon.HP > 20, increment count
	// Hint: Return the count

	return 0 // Replace with your implementation
}

// TODO TASK 7: Check if you have any fire type Pokemon
// Return true if at least one Pokemon is fire type
func hasFireType(team [3]Pokemon) bool {
	// YOUR CODE HERE
	// Hint: Loop through each Pokemon
	// Hint: If pokemon.Type == "fire", return true immediately
	// Hint: If loop finishes without finding fire, return false

	return false // Replace with your implementation
}

// TODO TASK 8: Find a water type Pokemon in your team
// Return the water Pokemon if found, otherwise return an empty Pokemon
func findWaterType(team [3]Pokemon) Pokemon {
	// YOUR CODE HERE
	// Hint: Loop through each Pokemon
	// Hint: If pokemon.Type == "water", return that Pokemon
	// Hint: If no water type found, return Pokemon{} (empty)

	return Pokemon{} // Replace with your implementation
}

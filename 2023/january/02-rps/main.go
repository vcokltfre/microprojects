package main

import (
	"fmt"
	"math/rand"
	"time"
)

var defaults = []string{
	"Rock",
	"Paper",
	"Scissors",
}

func generateWinMapping(objects []string) map[string]string {
	winMapping := map[string]string{}
	winMapping[objects[0]] = objects[len(objects)-1]

	for i := 1; i < len(objects); i++ {
		winMapping[objects[i]] = objects[i-1]
	}

	return winMapping
}

func getWinner(objects []string, p1, p2 string) string {
	if p1 == p2 {
		return "Draw"
	}

	winMapping := generateWinMapping(objects)

	if winMapping[p1] == p2 {
		return "Player 1 wins"
	}

	if winMapping[p2] == p1 {
		return "Player 2 wins"
	}

	return "Draw"
}

func isValidChoice(objects []string, choice string) bool {
	for _, o := range objects {
		if o == choice {
			return true
		}
	}
	return false
}

func runGame(objects ...[]string) {
	var choices []string
	if len(objects) > 0 {
		choices = objects[0]
	} else {
		choices = defaults
	}

	var p1, p2 string
	for {
		fmt.Print("Player 1: ")
		fmt.Scanln(&p1)
		if isValidChoice(choices, p1) {
			break
		}
		fmt.Println("Invalid choice")
	}

	rand.Seed(time.Now().UnixNano())
	p2 = choices[rand.Intn(len(choices))]

	fmt.Println(getWinner(choices, p1, p2))
}

func main() {
	fmt.Println("Default RPS:")
	runGame()

	fmt.Println("\nCustom RPS:")
	runGame([]string{
		"Rock",
		"Paper",
		"Scissors",
		"Lizard",
		"Spock",
	})
}

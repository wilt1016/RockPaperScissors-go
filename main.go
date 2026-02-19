package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

var wins = map[string]string{
	"rock":     "scissors",
	"paper":    "rock",
	"scissors": "paper",
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func determineWinner(user, computer string) string {
	if user == computer {
		return "tie"
	}
	if wins[user] == computer {
		return "win"
	}
	return "lose"
}

func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func pause(reader *bufio.Reader) {
	fmt.Println("Press enter to continue...")
	reader.ReadString('\n')
}

func isValidChoice(user string, choices []string) bool {
	for _, c := range choices {
		if user == c {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	choices := []string{"rock", "paper", "scissors"}

	for {
		clearScreen()

		fmt.Println("Pick rock, paper or scissors (or 'quit' to exit)")

		user, err := readInput(reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		user = strings.ToLower(user)

		if user == "quit" {
			break
		}

		computer := choices[rand.IntN(len(choices))]

		if !isValidChoice(user, choices) {
			fmt.Println("Invalid choice")
			pause(reader)
			continue
		}

		result := determineWinner(user, computer)
		switch result {
		case "win":
			fmt.Println("You win!")
			fmt.Printf("You chose %s, and the computer chose %s\n", user, computer)
		case "lose":
			fmt.Println("You lose!")
		case "tie":
			fmt.Println("Its a tie")

		}

		pause(reader)
		clearScreen()
	}
}

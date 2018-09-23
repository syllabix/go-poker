package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/syllabix/psychic-poker-player/player"
)

const (
	banner = `
  _____       ___         __
 / ___/___   / _ \ ___   / /__ ___  ____
/ (_ // _ \ / ___// _ \ /  '_// -_)/ __/
\___/ \___//_/    \___//_/\_\ \__//_/
	`

	errorHeader = `
ERROR…`
)

func main() {

	fmt.Printf("%s\n", banner)
	fmt.Println("Welcome to Go Poker - a super simple poker hand analyzer")
	fmt.Printf("Follow the prompts to begin. Type 'Exit' and press enter if you want to quit\n\n")
	prompt()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}

		result, err := player.RevealBestHand(input)

		if err != nil {
			printError(err)
		} else {
			fmt.Println()
			fmt.Println("Result:")
			fmt.Println(result)
			fmt.Println()
		}
		prompt()
	}

	if scanner.Err() != nil {
		printError(scanner.Err())
		os.Exit(1)
	}

	fmt.Println("Thanks for playing!")
}

func printError(err error) {
	fmt.Println(errorHeader)
	fmt.Println(err)
	fmt.Println()
}

func prompt() {
	fmt.Println("Type or paste your input and press enter:")
}

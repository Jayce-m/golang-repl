package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("POKEDEX > ")
		scanner.Scan()

		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}

		// The first string in the input is the command
		command := cleaned[0]

		value, exists := getCommands()[command] // This is probably not great..
		if exists {
			value.callback()
		} else {
			fmt.Println("Sorry thats not a valid command")
		}

	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location available.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations.",
			callback:    commandMapBack,
		},
	}
}

func cleanInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}

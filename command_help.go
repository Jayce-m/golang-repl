package main

import "fmt"

// Prints info about available commands
func commandHelp() {
	fmt.Println("AVAILABLE COMMANDS:")

	commandsAvailable := getCommands()

	for _, value := range commandsAvailable {
		fmt.Println("\t" + value.name + " - " + value.description)
	}
}

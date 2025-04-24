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
	callback    func(...string) error
}

func repl() {
	var param []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("notes > ")
		scanner.Scan()
		textBytes := scanner.Text()
		cleanText := cleanInput(string(textBytes))
		if len(cleanText) == 0 {
			continue
		}
		firstWord := cleanText[0]

		if len(cleanText) > 1 {
			param = cleanText[0:]
		}
		command, ok := getCommands()[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(param...)
		if err != nil {
			panic(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the gonotes",
			callback:    commandExit,
		},
	}
}

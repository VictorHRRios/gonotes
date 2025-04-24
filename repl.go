package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform not support clear screen")
	}
}

func repl() {
	var param []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		callClear()
		now := time.Now()
		fmt.Printf(
			`--------------------------------------------------
		%v
--------------------------------------------------
notes > `, now.Format(time.DateTime))
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

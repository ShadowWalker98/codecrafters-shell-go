package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	validCommands := map[string]bool{}
	validCommands["exit"] = true
	validCommands["echo"] = true
	for {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error encountered while receiving input")
		}
		command = command[:len(command)-1]
		commandSplit := strings.Split(command, " ")

		firstCommand := strings.TrimRight(commandSplit[0], "\n")

		// check if the command is a valid command
		if _, ok := validCommands[firstCommand]; ok {
			if firstCommand == "exit" {
				os.Exit(0)
			} else if firstCommand == "echo" {
				_, err = fmt.Fprint(os.Stdout, strings.Join(commandSplit[1:], " ")+"\n")

			}
		} else {
			_, err = fmt.Fprint(os.Stdout, command+": command not found\n")
		}
	}
}

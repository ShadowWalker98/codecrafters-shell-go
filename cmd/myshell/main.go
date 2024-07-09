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
	validCommands["type"] = true
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
			} else if firstCommand == "type" {
				typeArg1 := strings.TrimRight(commandSplit[1], "\n")
				if _, present := validCommands[typeArg1]; present {
					_, err = fmt.Fprint(os.Stdout, typeArg1+" is a shell builtin\n")
				} else {
					handlePathQueries(typeArg1)
				}
			}
		} else {
			_, err = fmt.Fprint(os.Stdout, command+": command not found\n")
		}
	}
}

func handlePathQueries(typeArg1 string) {
	env := os.Getenv("PATH")
	paths := strings.Split(env, ":")

	for _, path := range paths {
		exec := path + "/" + path
		if _, err := os.Stat(exec); err == nil {
			_, err := fmt.Fprintf(os.Stdout, "%v is %v\n", typeArg1, exec)
			if err != nil {
				return
			}
			return
		}
	}
	_, err := fmt.Fprint(os.Stdout, typeArg1+": not found\n")
}

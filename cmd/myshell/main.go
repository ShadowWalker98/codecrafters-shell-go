package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	for {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error encountered while receiving input")
		}
		command = command[:len(command)-1]
		commandSplit := strings.Split(command, " ")
		if strings.TrimRight(commandSplit[0], "\n") == "exit" {
			os.Exit(0)
		}
		_, err = fmt.Fprint(os.Stdout, command+": command not found\n")
		if err != nil {
			return
		}
	}
}

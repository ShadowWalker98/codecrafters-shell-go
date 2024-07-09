package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	for {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error encountered while receiving input")
		}
		_, err = fmt.Fprint(os.Stdout, command[:len(command)-1]+": command not found\n")
		if err != nil {
			return
		}
	}
}

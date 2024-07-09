package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	//Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Error encountered while receiving input")
	}
	command = strings.Replace(command, "\r\n", "", -1)
	command = command + ": command not found"
	_, err = fmt.Fprint(os.Stdout, command)
	if err != nil {
		return
	}

}

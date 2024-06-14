package main

import (
	"bufio"
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	// fmt.Fprint(os.Stdout, "$ ")	
		
	running := true
	// Wait for user input
	for running {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Fields(input)
		if input == "exit 0\n" {
			running = false
		} else if args[0] == "echo" {
			fmt.Fprint(os.Stdout, strings.Join(args[1:], " ") + "\n")
		} else {
			fmt.Fprint(os.Stdout, input[:len(input) - 1] + ": command not found\n")	
		}
		
	}

}

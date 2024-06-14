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
	commands := []string{"echo", "type", "exit"}
	// Wait for user input
	for running {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Fields(input)
		if input == "exit 0\n" {
			running = false
		} else if args[0] == "echo" {
			fmt.Fprint(os.Stdout, strings.Join(args[1:], " ") + "\n")
		} else if args[0] == "type" {
			found := false
			for _, command := range commands {
				if args[1] == command {
					fmt.Fprint(os.Stdout, command + " is a shell builtin\n")
					found = true
				}
			}
			if !found {
				fmt.Fprint(os.Stdout, args[1] + ": not found\n")
			}
		} else {
			fmt.Fprint(os.Stdout, input[:len(input) - 1] + ": command not found\n")	
		}
		
	}

}

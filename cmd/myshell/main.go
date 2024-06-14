package main

import (
	"bufio"
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	// fmt.Fprint(os.Stdout, "$ ")	
		
	var running = true
	// Wait for user input
	for running {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if input == "exit 0\n" {
			running = false
		} else {
			fmt.Fprint(os.Stdout, input[:len(input) - 1] + ": command not found\n")	
		}
		
	}

}

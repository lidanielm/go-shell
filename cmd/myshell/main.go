package main

import (
	"bufio"
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
	"strings"
	"lookup"
)

func main() {
	// Uncomment this block to pass the first stage
	// fmt.Fprint(os.Stdout, "$ ")	
	
	path := os.Args[0]
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
			// Check if the command is a shell builtin
			for _, command := range commands {
				if args[1] == command {
					fmt.Fprint(os.Stdout, command + " is a shell builtin\n")
					found = true
				}
			}
			// Check if the file exists
			if LookPath(args[1]) != "" {
				fmt.Fprint(os.Stdout, args[1] + " is " + path + "\n")
				found = true
			}
			if !found {
				fmt.Fprint(os.Stdout, args[1] + ": not found\n")
			}
		} else {
			fmt.Fprint(os.Stdout, input[:len(input) - 1] + ": command not found\n")	
		}
		
	}

}

func fileExists(path string, cmd string) bool {
	files, _ := os.ReadDir(path)
	for _, file := range files {
		if file.Name() == cmd {
			return true
		}
		if file.IsDir() {
			return fileExists(path + "/" + file.Name(), cmd)
		}
	}
	return false
}
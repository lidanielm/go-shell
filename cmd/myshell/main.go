package main

import (
	"bufio"
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
	"strings"
	"os/exec"
)

func main() {
	// Uncomment this block to pass the first stage
	// fmt.Fprint(os.Stdout, "$ ")	
	
	running := true
	commands := []string{"echo", "type", "exit"}
	homeDir := getWD()
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
					break
				}
			}

			// Check if the file exists
			file, _ := exec.LookPath(args[1])
			if file != "" && !found {
				fmt.Fprint(os.Stdout, args[1] + " is " + file + "\n")
				found = true
			}

			// If not found, print error message
			if !found {
				fmt.Fprint(os.Stdout, args[1] + ": not found\n")
				continue
			}
		} else if args[0] == "pwd" {
			dir := getWD()
			fmt.Fprint(os.Stdout, dir + "\n")
		} else if args[0] == "cd" {
			// Absolute
			if args[1][0] == '/' {
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Fprint(os.Stdout, "cd: " + args[1] + ": No such file or directory\n")	
				}
			} else if args[1] == "~" {
				os.Chdir(homeDir)
			} else {
				fmt.Fprint(os.Stdout, "cd: " + args[1] + ": No such file or directory\n")
			}
		} else {
			// Executable
			file, _ := exec.LookPath(args[0])
			if file != "" {
				cmd := exec.Command(args[0], args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				continue
			} else {
				fmt.Fprint(os.Stdout, args[0] + ": not found\n")
			}
		}
	}
}

func getWD() string {
	dir, _ := os.Getwd()
	return dir
}
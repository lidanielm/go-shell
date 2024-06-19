package main

import (
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
	"strings"
	"os/exec"
	"path/filepath"
	"errors"
	"bufio"
	// "github.com/eiannone/keyboard"
	// "os/signal"
	"sort"
	// input_autocomplete "github.com/JoaoDanielRufino/go-input-autocomplete"
	// "github.com/azul3d/keyboard"
	// termbox "github.com/nsf/termbox-go"
)

func main() {
	running := true
	// history := []string{}
	// Wait for user input
	for running {
		fmt.Fprint(os.Stdout, os.Getenv("PWD") + " ~ $ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Fields(input)
		if input == "exit 0" {
			running = false
		} else if args[0] == "echo" {
			echoCmd(args)
		} else if args[0] == "type" {
			typeCmd(args)
		} else if args[0] == "pwd" {
			pwdCmd()
		} else if args[0] == "cd" {
			cdCmd(args)
		} else if args[0] == "ls" {
			lsCmd()
		} else {
			execCmd(args)
		}
	}
}

func echoCmd(args []string) {
	fmt.Fprint(os.Stdout, strings.Join(args[1:], " ") + "\n")
}

func pwdCmd() {
	fmt.Fprint(os.Stdout, os.Getenv("PWD") + "\n")
}

func typeCmd(args []string) {
	commands := []string{"echo", "type", "pwd", "cd"}
	// Check if the command is a shell builtin
	for _, command := range commands {
		if args[1] == command {
			fmt.Fprint(os.Stdout, command + " is a shell builtin\n")
			return
		}
	}

	// Check if the file exists
	file, _ := exec.LookPath(args[1])
	if file != "" {
		fmt.Fprint(os.Stdout, args[1] + " is " + file + "\n")
		return
	}

	// If not found, print error message
	fmt.Fprint(os.Stdout, args[1] + ": not found\n")
}

func searchCMDPrefix(prefix string) []string {
	commands := []string{}
	if strings.HasPrefix("echo", prefix) {
		commands = append(commands, "echo")
	}
	if strings.HasPrefix("type", prefix) {
		commands = append(commands, "type")
	}
	if strings.HasPrefix("pwd", prefix) {
		commands = append(commands, "pwd")
	}
	if strings.HasPrefix("cd", prefix) {
		commands = append(commands, "cd")
	}

	// Check if the file exists
	file, _ := exec.LookPath(prefix)
	if file != "" {
		commands = append(commands, file)
	}

	sort.Strings(commands)
	return commands
}

func cdCmd(args []string) {
	targetPath := args[1]
	if targetPath == "~" {
		targetPath = os.Getenv("HOME")
		os.Setenv("PWD", targetPath)
		return
	}
	if targetPath[0] != '/' {
		targetPath = filepath.Join(os.Getenv("PWD"), targetPath)
		os.Setenv("PWD", targetPath)
		return
	}
	if _, err := os.Stat(targetPath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s: No such file or directory\n", targetPath)
	}
}

func execCmd(args []string) {
	// Executable
	file, _ := exec.LookPath(args[0])
	if file != "" {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		fmt.Fprint(os.Stdout, args[0] + ": not found\n")
	}
}

func lsCmd() {
	// Get current directory
	dir := os.Getenv("PWD")
	// Open directory
	d, _ := os.Open(dir)
	// Read directory
	files, _ := d.Readdir(-1)
	// Print directory content
	for _, file := range files {
		fmt.Fprintln(os.Stdout, file.Name())
	}
}
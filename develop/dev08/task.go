package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	input = strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r")

	commands := strings.Split(input, "|")

	commandsArgs := make(map[string][]string)
	for _, item := range commands {
		args := strings.Split(strings.TrimSpace(item), " ")
		commandsArgs[args[0]] = args[1:]
	}
	args := strings.Split(input, " ")
	for command, arg := range commandsArgs {
		switch command {
		case "cd":
			if len(args) < 2 {
				return ErrNoPath
			}
			return os.Chdir(arg[0])
		case "pwd":
			currentPath, err := os.Getwd()
			fmt.Println(currentPath)
			return err
		case "exit":
			os.Exit(0)
		}
		cmd := exec.Command(command, arg...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

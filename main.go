package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/peterh/liner"
)

// Aliases for shell commands
var (
	aliases = map[string]string{
		"ll": "ls -l", // 'll' is an alias for 'ls -l'
	}
	history = []string{} // History of executed commands
)

func main() {
	// Create a new line reader for interactive input
	line := liner.NewLiner()
	defer line.Close()

	// Get the hostname and user for the shell prompt
	hostname, _ := os.Hostname()
	user := os.Getenv("USER")
	if user == "" {
		user = os.Getenv("USERNAME")
	}

	for {
		// Get the current working directory
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// Construct the shell prompt
		prompt := fmt.Sprintf("%s@%s:%s$ ", user, hostname, dir)
		// Read user input with the prompt
		input, err := line.Prompt(prompt)
		if err != nil {
			if err.Error() == "prompt canceled" {
				fmt.Println("Prompt canceled")
				return
			}
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// Exit the shell if 'exit' is entered
		if input == "exit" {
			os.Exit(0)
		}

		// Display command history if 'history' is entered
		if input == "history" {
			for i, h := range history {
				fmt.Printf("%d: %s\n", i+1, h)
			}
			continue
		}

		// Clear the terminal screen if 'clear' is entered
		if input == "clear" {
			fmt.Print("\033[H\033[2J")
			continue
		}

		// Execute the input command
		if err := execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Append the input to history
		history = append(history, input)
		line.AppendHistory(input)
	}
}

// execInput executes the given input command
func execInput(input string) error {
	// Trim whitespace and split input into arguments
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")

	// Check if the command is an alias and expand it
	if cmd, ok := aliases[args[0]]; ok {
		args = append(strings.Split(cmd, " "), args[1:]...)
	}

	switch args[0] {
	// Handle 'cd' command to change directories
	case "cd":
		if len(args) < 2 {
			return errors.New("path required for 'cd' command")
		}
		return os.Chdir(args[1])
	// Handle 'exit' command to terminate the shell
	case "exit":
		os.Exit(0)
	}

	// Prepare command execution based on OS
	var cmd *exec.Cmd
	if isWindows() {
		if args[0] == "ls" {
			args[0] = "dir" // Convert 'ls' to 'dir' for Windows
		}
		cmd = exec.Command("cmd", "/C", strings.Join(args, " ")) // Windows command execution
	} else {
		cmd = exec.Command(args[0], args[1:]...) // Unix-like command execution
	}

	// Set command output to standard output and error
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run() // Run the command
}

// isWindows checks if the current OS is Windows
func isWindows() bool {
	return strings.Contains(strings.ToLower(os.Getenv("OS")), "windows")
}

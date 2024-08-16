# Go Shell

This is a simple interactive shell implemented in Go. It supports executing standard shell commands, maintains a history of executed commands, and provides aliasing for certain commands. You can use it as a base for building more intricate shell implementation.

## Features

- **Command Execution**: Execute standard shell commands on Unix-like and Windows systems.
- **Command History**: View a history of previously executed commands using the `history` command.
- **Aliases**: Supports aliases for commands, such as `ll` for `ls -l`.
- **Change Directory**: Use the `cd` command to change the current directory.
- **Shell Prompt**: Displays a dynamic prompt with the user, hostname, and current directory.
- **Clear Screen**: Clear the terminal screen with the `clear` command.
- **Exit**: Exit the shell using the `exit` command.
- **Platform Support**: Automatically adjusts command execution for Unix-like and Windows systems.

## Installation

To build and run the shell, you need to have Go installed on your machine.

1. Clone this repository:
   ```bash
   git clone https://github.com/interimme/go-shell.git
   ```

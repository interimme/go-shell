# Go Shell

This is a simple interactive shell implemented in Go. It supports executing standard shell commands, maintains a history of executed commands, and provides aliasing for certain commands. You can use it as a base for building more intricate shell implementation.

## Features

- **Command Execution**: Execute basic shell commands like `ls`, `cd`, `pwd`, and more.
- **Command History**: View a history of previously executed commands using the `history` command.
- **Cross-Platform Compatibility**: Automatically adjusts command execution for Unix-like and Windows systems.
- **Custom Commands**: Extend the shell with custom Go commands.
- **Clear Screen**: Clear the terminal screen with the `clear` command.
- **Exit**: Exit the shell using the `exit` command.
- **Error Handling**: Basic error handling for unsupported commands or incorrect usage.

## Installation

To install and run the Go Shell application:

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/interimme/go-shell.git
    cd go-shell
    ```

2. **Build the Application**:
    ```bash
    go build -o go-shell
    ```

3. **Run the Application**:
    ```bash
    ./go-shell
    ```

## Usage

After running the `go-shell` executable, you can start typing commands as you would in a regular shell. Some basic commands are supported by default:

- `pwd` - Prints the current working directory.
- `cd <directory>` - Changes the current working directory.
- `ls` - Lists files in the current directory.

You can extend the shell by modifying the Go code and adding your own custom commands.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

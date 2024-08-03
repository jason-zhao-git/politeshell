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
var pleaseSaid = false

func execInput(input string) error {
    // Remove the newline character.
    input = strings.TrimSuffix(input, "\n")

    // Split the input separate the command and the arguments.
    args := strings.Split(input, " ")

	//gotto say please in the beginning
	if !pleaseSaid && args[0] != "Please" {
		fmt.Print("Hey be polite!\n")
		return nil
	} 
	
	if args[0] == "Please" {
		pleaseSaid = true
		args = args[1:] // Remove "Please" from the arguments
	}

	//TODO: mem the please if already said
	if len(args) < 1 {
		fmt.Print("What?\n")
		return nil
	}

    // Check for built-in commands.
    switch args[0] {
	case "cd":
		// Check if the path is provided
		if len(args) < 2 {
			return ErrNoPath
		}
		// Change the directory and return the error
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)
	
	pleaseSaid = false
    // Set the correct output device.
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout

    // Execute the command and return the error.
    return cmd.Run()
}
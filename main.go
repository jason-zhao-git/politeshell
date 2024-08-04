package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
)

var workingDir string

func main() {
	var err error

	// Initialize workingDir with the current working directory
	workingDir, err = os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting current directory:", err)
		return
	}
	fmt.Println("Initial Working Directory:", workingDir)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
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
var creeperCalled = false
var creeperCountDown = 0
// List of polite words
var politeWords = []string{
	"please",
	"plz",
	"pls",
	"plzz",
	"could you",
	"would you kindly",
	"kindly",
	"if you could",
	"could you please",
	"may i",
	"could i",
	"if you wouldn't mind",
	"would you mind",
	"can you please",
	"could you kindly",
	"would you be so kind",
	"please do",
	"plz do",
	"it would be great if you",
	"i would appreciate it if you could",
	"would you please",
}

func inList(word string, list []string) bool {
	lowerWord := strings.ToLower(word)
	for _, item := range list {
		if strings.ToLower(item) == lowerWord {
			return true
		}
	}
	return false
}

// Function to check if input starts with any polite phrase
func startsWithPolitePhrase(input string, list []string) (bool, int) {
	lowerInput := strings.ToLower(input)

	longestMatchLength := 0

	for _, phrase := range list {
		if strings.HasPrefix(lowerInput, phrase) {
			phraseLength := len(phrase)
			if phraseLength > longestMatchLength {
				longestMatchLength = phraseLength
			}
		}
	}
	return longestMatchLength > 0, longestMatchLength
}

// print creeper
func printFile(fileName string) error {
	// Use filepath.Join to construct the full path to creeper.txt
	artPath := filepath.Join(workingDir, "static", fileName)

	art, err := os.ReadFile(artPath)
	if err != nil {
		return err
	}
	fmt.Println(string(art))
	return nil
}

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	match, length := startsWithPolitePhrase(input, politeWords)
	//gotto say please in the beginning
	if !pleaseSaid && !match {
		fmt.Print("Hey be polite!\n")
		return nil
	}

	if match {
		pleaseSaid = true
		input = strings.TrimSpace(input[length:])
	}

	// Split the input with white spaces
	args := strings.Fields(input)

	if creeperCountDown >= 3 {
		fmt.Println("Boooom! Creeper Exploded gotta run 'cat' cmd to scare it away")
		printFile("explosion.txt")
		os.Exit(0)
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
		pleaseSaid = false
		// Change the directory and return the error
		return os.Chdir(args[1])

	case "creeper":
		pleaseSaid = false
		creeperCalled = true
		return printFile("creeper.txt")

	case "sword":
		pleaseSaid = false
		if creeperCalled {
			creeperCalled = false
			creeperCountDown = 0
			fmt.Print("Creeper is slayed")
			return printFile("defeatCreeper.txt")
		}
		
		return printFile("sword.txt")

	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	if args[0] == "cat" && creeperCalled {
		creeperCalled = false
		creeperCountDown = 0
		fmt.Print("Creeper is scared away!\n")
	}

	// after cmd exec, add one to creeperCount down
	if creeperCalled {
		creeperCountDown = creeperCountDown + 1
	}

	pleaseSaid = false
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}

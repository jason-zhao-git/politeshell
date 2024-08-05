# Polite Shell

Well, turns out a basic shell is not hard to build. Welcome to **Polite Shell**, an interactive command-line program that enforces politeness before executing commands. This shell adds a fun twist to your typical command-line experience by requiring polite language and offering unique commands.

## Features

- **Interactive Command Line:** Execute standard shell commands interactively in a Unix like system.
- **Politeness Enforcement:** Requires polite phrases before executing commands, with responses for impoliteness.
- **Command History:** Browse your input history using the up/down arrow keys.
- **Custom Commands:** Use special commands to affect shell's behavior.
- **Custom ASCII Art Display:** Displays ASCII art for specific commands.

## Installation

To use Polite Shell, ensure you have Go installed on your system. You can download Go from the official [Go website](https://golang.org/dl/).

### Steps to Install and Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/jason-zhao-git/politeshell.git
2. **Navigate to the project directory:**

   ```bash
   cd politeshell
3. **Run the project:**
   we could
   ```bash
   go build shell.go
   ./shell.go
   ```
   or
   ```bash
   go run shell.go
   ```

## Usage

Once you start Polite Shell, you'll see the prompt:

```bash
>>>
```
Here's how you can interact with Polite Shell:

Type shell commands as you normally would. To successfully execute a command, prepend it with a polite phrase like "please".
Example:
   ```bash
   >>> please
   >>> ls
   ```

   ```bash
   >>> please ls
   >>> plz do cd
   >>> would you kindly vim
   ```
Try some special custom commands:
Example: 
   ```bash
   >>> plz creeper
   >>> if you wouldn't mind sword
   >>> plzz cat fileName.txt
   ```
Exit shell using 'exit' or "^C"

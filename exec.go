package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Before executing input, need to parse input for quotes, etc.
// Parse rules regarding special characters are copied from
// bash.



// execInput executes command, return eventual error
// Checks for builtins
func execInput(input string, history *os.File) error {
	history.WriteString(input)
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	// Check for builtins
	switch args[0] {
		// Shall we exit?
		case "exit":
			os.Exit(0)
		// Change working directory
		case "cd":
			var dir string
			if len(args) < 2 {
				dir = home
			} else {
				dir = args[1]
			}
			if err := cd(dir); err != nil {
				fmt.Fprintln(os.Stderr, err)
		}
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

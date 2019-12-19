package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	// Check for builtins
	switch args[0] {
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

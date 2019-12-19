package main

import (
	"bufio"
	"fmt"
	"os"
)

var version string
var home string

func main(){
	version = "0.0.1"
	fmt.Println("Welcome to the gshell.")
	fmt.Println("gshell version: " + version)
	home, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(home + "$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

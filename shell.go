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
	history, err := os.Create("/tmp/dat2")
	if err != nil {
		fmt.Fprintln((os.Stderr, err))
	}
	defer history.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(home + "$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input, history); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

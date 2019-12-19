package main

import "os"

func cd(path string) error {
	return os.Chdir(path)
}

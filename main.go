package main

import (
	"fmt"
	"os"

	"github.com/skanehira/diff-file/gui"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}

	file := args[1]
	if err := gui.New(file).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

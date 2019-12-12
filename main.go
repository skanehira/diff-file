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

	if err := gui.New().Run(args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

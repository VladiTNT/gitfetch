package main

import (
	"fmt"
	"os"
)

func main() {
	cli := NewCli()
	if err := cli.Exec(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

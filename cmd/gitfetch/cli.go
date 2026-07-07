package main

import (
	"flag"
	"fmt"
)

type Cli struct {
	Dir      *string
	AvgChars *bool
}

func NewCli() *Cli {
	cli := &Cli{
		Dir:      flag.String("dir", ".", "directory"),
		AvgChars: flag.Bool("avgchars", false, "average amount of chars per line"),
	}

	flag.Parse()

	return cli
}

func (c *Cli) Exec() error {
	var files []*FileInfo

	// Get all of the files in the project
	if err := ParseDirectory(*c.Dir, &files); err != nil {
		return err
	}

	var totalSize int64
	var totalLines int

	for _, file := range files {
		totalSize += file.Size
		totalLines += file.Lines
	}

	fmt.Printf("Total size: %d\n", totalSize)
	fmt.Printf("Total lines: %d\n", totalLines)

	return nil
}

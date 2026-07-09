package main

import (
	"flag"
	"fmt"
)

type Cli struct {
	Dir      *string
	AvgChars *bool
	NoSize   *bool
	NoLines  *bool
}

func NewCli() *Cli {
	cli := &Cli{
		Dir:      flag.String("dir", ".", "directory"),
		AvgChars: flag.Bool("avgchars", false, "average amount of chars per line"),
		NoSize:   flag.Bool("nosize", false, "disable total project size"),
		NoLines:  flag.Bool("nolines", false, "disable total project lines"),
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

	fmt.Println("Project statistics:")

	if !*c.NoSize {
		fmt.Printf("Total project size: %d bytes.\n", GetTotalSize(files))
	}

	if !*c.NoLines {
		fmt.Printf("Total lines: %d.\n", GetTotalLines(files))
	}

	if *c.AvgChars {
		fmt.Printf("Average line length: %.2f.\n", GetTotalAvgChars(files))
	}

	return nil
}

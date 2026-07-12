package main

import (
	"flag"
	"fmt"
)

type Cli struct {
	Dir      *string
	NoSize   *bool
	NoLines  *bool
	Nolang   *bool
	AvgChars *bool

	DefaultIgnoredEntries        []string
	DefaultIgnoredFileExtensions []string
}

func NewCli() *Cli {
	cli := &Cli{
		Dir:      flag.String("dir", ".", "directory"),
		NoSize:   flag.Bool("nosize", false, "disable total project size"),
		NoLines:  flag.Bool("nolines", false, "disable total project lines"),
		Nolang:   flag.Bool("noalng", false, "disable language makeup"),
		AvgChars: flag.Bool("avgchars", false, "average amount of chars per line"),

		DefaultIgnoredEntries: []string{
			".git", ".gitignore",
		},
		DefaultIgnoredFileExtensions: []string{},
	}

	flag.Parse()

	return cli
}

func (c *Cli) Exec() error {
	var files []*FileInfo

	// Get all of the files in the project
	if err := ParseDirectory(*c.Dir, &files, c.DefaultIgnoredEntries); err != nil {
		return err
	}

	fmt.Println("Project statistics:")

	if !*c.NoSize {
		fmt.Printf("Total project size: %d bytes.\n", GetTotalSize(files))
	}

	if !*c.NoLines {
		fmt.Printf("Total lines: %d.\n", GetTotalLines(files))
	}

	if !*c.Nolang {
		langs, size := GetLanguageMakeup(files)

		for k, v := range langs {
			fmt.Printf("%s files are %.2f%%.\n", k, float64(v)/float64(size)*100)
		}
	}

	if *c.AvgChars {
		fmt.Printf("Average line length: %.2f.\n", GetTotalAvgChars(files))
	}

	return nil
}

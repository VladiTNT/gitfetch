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
	IgnEnt   *string
	IgnExt   *string
	NoDefEnt *bool
	NoDefExt *bool
	NoDefAss *bool

	Ignores Ignores
}

func NewCli() *Cli {
	cli := &Cli{
		Dir:      flag.String("root", ".", "project root"),
		NoSize:   flag.Bool("nosz", false, "disable total project size"),
		NoLines:  flag.Bool("noln", false, "disable total project lines"),
		Nolang:   flag.Bool("nolang", false, "disable language makeup"),
		AvgChars: flag.Bool("avgchars", false, "average amount of chars per line"),
		IgnEnt:   flag.String("ignent", "", "add more file entry ignores"),
		IgnExt:   flag.String("ignext", "", "add more file extensions ignores"),
		NoDefEnt: flag.Bool("nodefent", false, "no default entry ignores"),
		NoDefExt: flag.Bool("nodefext", false, "no default extension ignores"),
		NoDefAss: flag.Bool("nodefass", false, "no default asset ignores"),

		Ignores: NewIgnores(),
	}

	flag.Parse()

	return cli
}

func (c *Cli) Exec() error {
	var files []*FileInfo

	// Making final ignored extensions
	finalIgnoredExtensions := append(c.Ignores.Extensions, c.Ignores.AssetExtensions...)

	// Get all of the files in the project
	if err := ParseDirectory(*c.Dir, &files, c.Ignores.Entries, finalIgnoredExtensions); err != nil {
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

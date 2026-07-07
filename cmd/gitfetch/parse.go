package main

import (
	"fmt"
	"os"
)

func ParseDirectory(name string) error {
	entries, err := os.ReadDir(name)
	if err != nil {
		return err
	}

	var files []*FileInfo

	for _, entry := range entries {
		if entry.IsDir() {
			if err := ParseDirectory(name + "/" + entry.Name()); err != nil {
				return err
			}
		} else {
			fileInfo, err := ParseFile(name + "/" + entry.Name())
			if err != nil {
				return err
			}

			files = append(files, fileInfo)
		}
	}

	var totalSize int64
	var totalLines int

	for _, file := range files {
		totalSize += file.Size
		totalLines += file.Lines
	}

	fmt.Printf("Size: %d, Lines: %d\n", totalSize, totalLines)

	return nil
}

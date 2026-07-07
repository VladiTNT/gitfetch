package main

import "fmt"

func main() {
	var files []*FileInfo
	if err := ParseDirectory(".", &files); err != nil {
		fmt.Printf("Error parsing directories: %v\n", err)
	}

	var totalSize int64
	var totalLines int

	for _, file := range files {
		totalSize += file.Size
		totalLines += file.Lines
	}

	fmt.Printf("Size: %d, Lines: %d\n", totalSize, totalLines)
}

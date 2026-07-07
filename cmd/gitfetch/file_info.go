package main

import (
	"bufio"
	"os"
)

type FileInfo struct {
	Size     int64
	Lines    int
	AvgChars int
}

func ParseFile(name string) (*FileInfo, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fileInfo := new(FileInfo)
	fileInfo.Size = info.Size()

	sc := bufio.NewScanner(f)

	// total amount of characters in the file
	var totalCharAmount int

	for sc.Scan() {
		totalCharAmount += len(sc.Text())

		fileInfo.Lines++
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	// we divide total amount of characters by the amount of lines to get the
	// average amount of characters in each line of code
	fileInfo.AvgChars = totalCharAmount / fileInfo.Lines

	return fileInfo, nil
}

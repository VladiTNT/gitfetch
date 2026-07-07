package main

import (
	"bufio"
	"os"
)

type FileInfo struct {
	Size  int64
	Lines int
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

	for sc.Scan() {
		fileInfo.Lines++
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return fileInfo, nil
}

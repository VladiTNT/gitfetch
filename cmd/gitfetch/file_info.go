package main

import (
	"bufio"
	"os"

	"github.com/VladiTNT/gitfetch/pkg/files"
)

type FileInfo struct {
	Size      int64
	Lines     int
	AvgChars  float64
	Extension string
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
	fileInfo.AvgChars = float64(totalCharAmount) / float64(fileInfo.Lines)

	fileInfo.Extension = files.GetFileExtension(name)

	return fileInfo, nil
}

func GetTotalSize(files []*FileInfo) int64 {
	var totalSize int64

	for _, file := range files {
		totalSize += file.Size
	}

	return totalSize
}

func GetTotalLines(files []*FileInfo) int {
	var totalLines int

	for _, file := range files {
		totalLines += file.Lines
	}

	return totalLines
}

func GetTotalAvgChars(files []*FileInfo) float64 {
	var totalAvgChars float64

	for _, file := range files {
		totalAvgChars += file.AvgChars
	}

	return totalAvgChars / float64(len(files))
}

func GetLanguageMakeup(files []*FileInfo) (map[string]int64, int64) {
	langs := make(map[string]int64)

	for _, file := range files {
		langs[file.Extension] += file.Size
	}

	var size int64

	for _, v := range langs {
		size += v
	}

	return langs, size
}

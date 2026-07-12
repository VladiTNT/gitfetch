package main

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

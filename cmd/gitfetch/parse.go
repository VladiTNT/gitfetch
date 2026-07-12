package main

import (
	"os"
	"slices"

	"github.com/VladiTNT/gitfetch/pkg/tools"
)

func ParseDirectory(name string, files *[]*FileInfo, ignores, ignoredExtensions []string) error {
	entries, err := os.ReadDir(name)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		// Ignore entries
		if slices.Contains(ignores, entry.Name()) {
			continue
		}

		if entry.IsDir() {
			if err := ParseDirectory(name+"/"+entry.Name(), files, ignores, ignoredExtensions); err != nil {
				return err
			}
		} else {
			// Ignore files with extensions that we don't want to count
			if slices.Contains(ignoredExtensions, tools.GetFileExtension(entry.Name())) {
				continue
			}

			fileInfo, err := ParseFile(name + "/" + entry.Name())
			if err != nil {
				return err
			}

			*files = append(*files, fileInfo)
		}
	}

	return nil
}

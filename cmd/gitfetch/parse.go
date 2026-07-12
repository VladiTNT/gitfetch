package main

import (
	"os"
	"slices"
)

func ParseDirectory(name string, files *[]*FileInfo, ignores []string) error {
	entries, err := os.ReadDir(name)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if slices.Contains(ignores, entry.Name()) {
			continue
		}

		if entry.IsDir() {
			if err := ParseDirectory(name+"/"+entry.Name(), files, ignores); err != nil {
				return err
			}
		} else {
			fileInfo, err := ParseFile(name + "/" + entry.Name())
			if err != nil {
				return err
			}

			*files = append(*files, fileInfo)
		}
	}

	return nil
}

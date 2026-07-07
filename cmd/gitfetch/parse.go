package main

import (
	"os"
)

func ParseDirectory(name string, files *[]*FileInfo) error {
	entries, err := os.ReadDir(name)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if err := ParseDirectory(name+"/"+entry.Name(), files); err != nil {
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

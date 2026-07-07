package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParseDirectory(name string) error {
	entries, err := os.ReadDir(name)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if err := ParseDirectory(name + "/" + entry.Name()); err != nil {
				return err
			}
		} else {
			if err := ParseFile(name + "/" + entry.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func ParseFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var n int
	for sc.Scan() {
		n++
	}

	if err := sc.Err(); err != nil {
		return err
	}

	fmt.Printf("File %s has %d lines.\n", name, n)

	return nil
}

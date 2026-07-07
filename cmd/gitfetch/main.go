package main

import "fmt"

func main() {
	if err := ParseDirectory("."); err != nil {
		fmt.Printf("Error parsing directories: %v\n", err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/VladiTNT/gitfetch/pkg/assets"
)

type Ignores struct {
	Entries         []string `json:"entries"`
	Extensions      []string `json:"extensions"`
	AssetExtensions []string `json:"assets"`
}

func NewIgnores() Ignores {
	var ignores Ignores

	if err := json.Unmarshal(assets.IgnoresJsonData, &ignores); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	return ignores
}

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	paths := []string{
		"E:\\project\\main.go",
		"E:\\project\\utils\\helper.go",
		"E:\\project\\readme.md",
		"E:\\project\\config.json",
	}

	for _, p := range paths {
		// Ambil nama file
		filename := filepath.Base(p)

		// Cek apakah .go
		match, err := filepath.Match("*.go", filename)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		if match {
			// Ambil folder (parent directory)
			dir := filepath.Base(filepath.Dir(p))

			fmt.Printf("File: %s | Folder: %s\n", filename, dir)
		}
	}
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	demoDir := "demo_files"

	err := os.MkdirAll(demoDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Folder dibuat:", demoDir)

	filePath := filepath.Join(demoDir, "catatan.txt")
	isi := "Halo dari Go!\nIni file yang dibuat oleh program.\nBaris ketiga."

	err = os.WriteFile(filePath, []byte(isi), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File ditulis:", filePath)
}

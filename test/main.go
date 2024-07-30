package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Error retrieving user home dir: %s", err)
	}

	filename := "BarcodePDF.pdf"
	desktopPath := filepath.Join(homeDir, "Desktop", filename)
	fmt.Println(desktopPath)
}

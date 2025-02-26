package main

import (
	"log"
	"path/filepath"
)

// TODO: replace with user input
const (
	defaultTemplate   = "default"
	defaultOutputPath = "."
)

func main() {
	// TODO: get go module name as an input -> run go mod init with the name
	// Command from the command line
	templatePath := filepath.Join("templates", defaultTemplate)
	// Copy selected template into target directory

	if err := CopyDir(templatePath, defaultOutputPath); err != nil {
		log.Printf("Error copying template: %v", err)
	}
}

// CopyDir copies a directory from src to dst recursively
func CopyDir(src, dst string) error {
	// copy

	return nil
}

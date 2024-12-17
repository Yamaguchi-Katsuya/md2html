package main

import (
	"fmt"
	"os"

	"github.com/gomarkdown/markdown"
)

func md2html(inputPath string, outputPath string) error {
	mdContent, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	htmlContent := markdown.ToHTML(mdContent, nil, nil)
	err = os.WriteFile(outputPath, htmlContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run md2html.go <input_path> <output_path>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", inputPath)
		os.Exit(1)
	}

	err := md2html(inputPath, outputPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("File %s converted to %s\n", inputPath, outputPath)
}

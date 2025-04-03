package main

import (
	"fmt"
	"go-reloaded/proces"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . <input_file.txt> [output_file.txt]")
		return
	}

	// Get input file path from command-line arguments
	inputFile := os.Args[1]

	// Check if the input file has a .txt extension
	if filepath.Ext(inputFile) != ".txt" {
		fmt.Println("Error: Input file must have a .txt extension.")
		return
	}

	// Determine the output file path
	var outputFile string
	if len(os.Args) == 3 {
		outputFile = os.Args[2]
	} else {
		// Default output file name if not provided
		outputFile = "result.txt"
	}

	// Ensure the output file has a .txt extension
	if filepath.Ext(outputFile) != ".txt" {
		fmt.Println("Error: Output file must have a .txt extension.")
		return
	}

	// Ensure input and output files are not the same
	if filepath.Clean(inputFile) == filepath.Clean(outputFile) {
		fmt.Println("Error: Input and output files must not be the same.")
		return
	}

	// Read the input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the input text into lines
	lines := strings.Split(string(data), "\n")

	// Process each line
	for i, textLine := range lines {
		textLine = proces.ProcessTrans(textLine) // Process uppercase conversion
		textLine = proces.ReplaceNums(textLine)
		textLine = proces.ProscessPuncQ(textLine)
		textLine = proces.Articl(textLine) // Process articles (commented out for now)
		lines[i] = textLine                // Update the line with processed text
	}

	// Join the processed lines back into a single string with newlines
	outputText := strings.Join(lines, "\n")

	// Write the processed text to the output file
	err = os.WriteFile(outputFile, []byte(outputText), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Text processing complete. Check the output file:", outputFile)
}

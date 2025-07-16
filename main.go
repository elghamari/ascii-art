package main

import (
	"fmt"
	"os"

	"ascii-art/functions"
)

func main() {
	// Check if exactly one argument is provided
	if len(os.Args) != 2 {
		fmt.Println("you must enter 1 argument")
		return
	}
	input := os.Args[1]
	myfile := "standard.txt"
	// Validate the input string if is printable and contains only valid characters
	if !functions.IsValidInput(input) {
		fmt.Println("Error: input contains invalid characters")
		return
	}

	lines := functions.ParseInput(input) // Split the input into lines based on "\n"
	data, err := os.ReadFile(myfile) // Read the font file
	if err != nil {
		fmt.Println("your fille have an error", err)
		return
	}
	fontTable := functions.BuildFontTable(string(data)) // Build the font table from file data
	functions.ProcessLines(lines, fontTable) // Convert and print ASCII art for each line
}

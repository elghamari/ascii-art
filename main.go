package main

import (
	"fmt"
	"os"

	"ascii-art/functions" 
)

// printUsage: helper function to show how to use the program
func printUsage() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println("Example: go run . something standard")
}

func main() {
	// Check if exactly two arguments are passed (excluding the program name)
	if len(os.Args) != 3 {
		printUsage()
		return
	}

	// Extract input string and banner name from command-line arguments
	input := os.Args[1]
	bannerName := os.Args[2]

	// Map of available banners and their corresponding file paths
	banners := map[string]string{
		"standard":   "banners/standard.txt",
		"shadow":     "banners/shadow.txt",
		"thinkertoy": "banners/thinkertoy.txt",
	}

	// Check if the banner name exists in the map
	bannerFile, ok := banners[bannerName]
	if !ok {
		printUsage()
		return
	}

	// Validate the input string if is printable and contains only valid characters
	if !functions.IsValidInput(input) {
		fmt.Println("Error: input contains invalid characters")
		return
	}

	// Split the input into lines based on "\n"
	lines := functions.ParseInput(input)

	// Read the banner file content
	data, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading banner file", err)
		return
	}

	
	fontTable := functions.BuildFontTable(string(data))  // Build the font table from file data

	functions.ProcessLines(lines, fontTable) // Convert and print ASCII art for each line
}

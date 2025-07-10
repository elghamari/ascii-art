package main

import (
	"fmt"
	"os"

	"ascii-art/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you must enter 1 argument")
		return
	}
	input := os.Args[1]
	myfile := "standard.txt"
	if !functions.IsValidInput(input) {
		fmt.Println("Error: input contains invalid characters")
		return
	}

	lines := functions.ParseInput(input)
	data, err := os.ReadFile(myfile)
	if err != nil {
		fmt.Println("your fille have an error", err)
		return
	}
	fontTable := functions.BuildFontTable(string(data))
	functions.ProcessLines(lines, fontTable)
}

package functions

import (
	"fmt"
	"strings"
)
// ProcessLines takes each line of input, converts it to ASCII art using the fontTable, and prints it.
// It handles empty lines by printing a blank line.
func ProcessLines(lines []string, fontTable map[int]string) {
	var finaleSlice [][]string
	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println()
			continue
		}
		for _, val := range line {
			for key, v := range fontTable {
				if key == int(val) {
					newslice := strings.Split(v, "\n")
					finaleSlice = append(finaleSlice, newslice)
				}
			}
		}
		PrintAsciiArt(finaleSlice)

		finaleSlice = nil
	}
}



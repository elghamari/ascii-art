package functions

import (
	"fmt"
	"strings"
)
// PrintAsciiArt prints the ASCII art representation line by line.
// It takes a 2D slice where each inner slice represents the ASCII art for a character.
func PrintAsciiArt(slice [][]string) {
	for i := 0; i < 8; i++ {
		myrows := []string{}
		for _, val := range slice {
			myrows = append(myrows, val[i])
		}
		fmt.Println(strings.Join(myrows, ""))
	}
}

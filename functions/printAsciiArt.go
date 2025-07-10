package functions

import (
	"fmt"
	"strings"
)

func PrintAsciiArt(slice [][]string) {
	for i := 0; i < 8; i++ {
		myrows := []string{}
		for _, val := range slice {
			myrows = append(myrows, val[i])
		}
		fmt.Println(strings.Join(myrows, ""))
	}
}

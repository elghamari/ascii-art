package functions

import "strings"

func BuildFontTable(data string) map[int]string {
	fontTable := make(map[int]string)
	chunks := strings.Split(data, "\n\n")

	if strings.Count(chunks[0], "\n") == 8 {
		chunks[0] = chunks[0][1:]
	}
	key := 32

	for _, val := range chunks {
		 fontTable[key] = val 
			key++
	}
	return fontTable
}
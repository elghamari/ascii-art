package functions

import "strings"


// ParseInput splits the input string by the literal "\n" and returns a slice of strings.
// It also removes a leading empty string if present.
func ParseInput(input string) []string {
	myinput := strings.Split(input, "\\n")
	for i, val := range myinput {
		if val != "" {
			break
		}
		if i == len(myinput)-1 && val == "" {
			myinput = myinput[1:]
		}
	}
	return myinput
}


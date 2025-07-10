package functions

import "strings"

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


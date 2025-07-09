package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you must enter 1 argument")
		return
	}
	input := os.Args[1]
	myfile := "standard.txt"

	data, err := os.ReadFile(myfile)
	if err != nil {
		fmt.Println("your fille have an error", err)
		return
	}
	myinput := strings.Split(input, "\\n")
	// newinput := strings.Join(myinput, " ")
	myData := strings.Split(string(data), "\n\n")

	 fmt.Println(myinput)
	// fmt.Println(myData)
	if strings.Count(string(myData[0]), "\n") == 8 {
		myData[0] = myData[0][1:]
	}
	myTables := make(map[int]string)
	key := 32
	for _, val := range myData {
		myTables[key] = val
		key++
	}
	// fmt.Println(myTables)
	var finaleSlice [][]string
	for _, value := range myinput {
		if len(value) == 0 {
			fmt.Println()
			continue
		}
		for _, val := range value {
			for key, valk := range myTables {
				if key == int(val) {
					newslice := strings.Split(valk, "\n")
					finaleSlice = append(finaleSlice, newslice)
				}
			}
		}
		Print(finaleSlice)

		finaleSlice = nil
	}
}

func Print(s [][]string) {
	for i := 0; i < 8; i++ {
		myrows := []string{}
		for _, val := range s {
			myrows = append(myrows, val[i])
		}
		fmt.Println(strings.Join(myrows, " "))

	}
}

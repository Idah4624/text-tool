package main

import (
	"fmt"
	"os"
	"strings"
	//"strings"
)


func main() {
	//fmt.Println("I am building a text tool")

	//input := os.Args[1]
	//output := os.Args[2]
	 if len(os.Args) != 3 {
		fmt.Println("usage:go run . input.txt output txt")
	 	return
	 }

	input := os.Args[2]
	output := os.Args[1]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error occured while reading file", err)
		return
	}

	value := strings.ToUpper(string(data))
	err = os.WriteFile(output, []byte(value), 0655)
	if err != nil {
		fmt.Println("error copying file", err)
		return
	}


}
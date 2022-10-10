package main

import (
	"fmt"
	"os"
	"strings"
)

// dividing string and counting words
func count(str string) int {
	str = strings.TrimRight(str, "\r\n")
	strArr := strings.Fields(str)

	return len(strArr)
}

// read incoming value from console
func readInput() string {
	input := ""
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	return input
}

func main() {
	str := readInput()
	result := count(str)

	fmt.Println(result)
}

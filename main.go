package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: at least one argument is required.")
		os.Exit(1)
	}

	input := os.Args[1]

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error converting input to int:", err)
	} else {
		output := numToWord(number)
		fmt.Printf("%d -> '%s'\n", number, output)
	}
}

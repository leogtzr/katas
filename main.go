package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	for {
		fmt.Printf("Enter some text: ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF { // EOF checking...
				fmt.Println("EOF detected. Exiting.")
				os.Exit(0)
			}
			fmt.Println("Error reading input:", err)
			return
		}

		input = input[:len(input)-1]

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error converting input to int:", err)
		} else {
			output := numToWord(number)
			fmt.Printf("%d -> '%s'\n", number, output)
		}
	}
}

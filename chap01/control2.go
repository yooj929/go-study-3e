package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}

	argument := os.Args[1]
	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int :", argument)
	}
	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negative integer")
	default:
		fmt.Println("This should not happen:", value)
	}
}

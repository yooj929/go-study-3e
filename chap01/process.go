package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	fmt.Println(arguments)
	var total, nInts, nFloats int
	invalid := make([]string, 0)

	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		invalid = append(invalid, k)
	}
	fmt.Println(total, nInts, nFloats)

}

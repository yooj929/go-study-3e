package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
	}
	data = append(data, Entry{"A", "B", "C"})
	data = append(data, Entry{"D", "E", "F"})
	data = append(data, Entry{"G", "H", "I"})
	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])

		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}

		fmt.Println(*result)

	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}

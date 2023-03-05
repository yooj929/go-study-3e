package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, direcotory := range pathSplit {
		fullPath := filepath.Join(direcotory, file)

		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()

			if mode.IsRegular() {
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}

}

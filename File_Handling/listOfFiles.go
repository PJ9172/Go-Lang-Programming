//Write a program that lists all the files in a directory with their size and modification date.

package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir("./")
	if err != nil {
		fmt.Println("Error to read directory!!!", err)
		return
	}

	for _, file := range files {
		info, _ := file.Info()
		fmt.Printf("%s\t %d\t %s\n", info.Name(), info.Size(), info.ModTime())
	}
}

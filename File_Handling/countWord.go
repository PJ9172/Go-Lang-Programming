//Write a program that reads a file and counts the number of words in it.

package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// func main() {
// 	file, err := os.Open("example.txt")
// 	if err != nil {
// 		fmt.Println("Error to open file!!!", err)
// 		return
// 	}
// 	defer file.Close()

// 	var str string
// 	buffer := make([]byte, 100)
// 	for {
// 		n, err := file.Read(buffer)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("Error to read file!!!", err)
// 			return
// 		}
// 		str = str + string(buffer[:n])
// 	}
// 	word := strings.Fields(str)
// 	fmt.Println("Count of words in file : ", len(word))
// }

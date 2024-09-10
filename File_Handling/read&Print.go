//Write a program that reads the contents of a file and prints it to the console.

package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("D:/Prajwal/GoLang/File_Handling/example.txt")
// 	if err != nil {
// 		fmt.Println("Error in file opening!!!", err)
// 		return
// 	}
// 	defer file.Close()
// 	buffer := make([]byte, 100)
// 	for {
// 		n, err := file.Read(buffer)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Print("Error in file reading!!!")
// 		}
// 		fmt.Print(string(buffer[:n]))
// 	}
// }

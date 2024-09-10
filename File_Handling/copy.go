//Write a program that reads the contents of one file and writes it to another file.

package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("newFile.txt")
// 	if err != nil {
// 		fmt.Println("Error to open file!!!", err)
// 		return
// 	}
// 	defer file.Close()

// 	wfile, err := os.Create("copy.txt")
// 	if err != nil {
// 		fmt.Println("Error to create file!!!", err)
// 		return
// 	}
// 	defer wfile.Close()

// 	buffer := make([]byte, 100)
// 	for {
// 		bytecount, err := file.Read(buffer)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("Error to read file!!!", err)
// 			return
// 		}
// 		_, err = wfile.WriteString(string(buffer[:bytecount]))
// 		if err != nil {
// 			fmt.Println("Error to write file!!", err)
// 			return
// 		}
// 	}
// 	fmt.Println("Copying files!!!")
// }

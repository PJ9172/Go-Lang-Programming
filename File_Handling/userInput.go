//Write a program that writes user input into a file, and then reads the file and prints the contents.

package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Print("Enter the content of file : ")
// 	reader := bufio.NewReader(os.Stdin)
// 	content, _, err := reader.ReadLine()
// 	if err != nil {
// 		fmt.Print("Error to read input!!", err)
// 		return
// 	}
// 	// fmt.Println("You enterd : ", string(content))

// 	file, err := os.Create("newFile.txt")
// 	if err != nil {
// 		fmt.Print("Error to create newfile!!!", err)
// 		return
// 	}
// 	defer file.Close()
// 	n, err := file.WriteString(string(content))
// 	if err != nil {
// 		fmt.Print("Error to write file!!!")
// 		return
// 	}
// 	fmt.Println("Write ", n, "bytes in file!!!")
// }

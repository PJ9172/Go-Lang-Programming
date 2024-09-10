//Write a program that appends new data to an existing file without overwriting the original content.

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
// 		fmt.Println("Error in reading content!!!", err)
// 		return
// 	}

// 	file, err := os.OpenFile("D:/Prajwal/GoLang/File_Handling/newFile.txt",os.O_APPEND|os.O_WRONLY,0644)
// 	if err != nil {
// 		fmt.Println("Error to open file!!", err)
// 		return
// 	}

// 	defer file.Close()

// 	n,err := file.WriteString("\n"+string(content))
// 	if err != nil {
// 		fmt.Println("Error to write file!!!", err)
// 		return
// 	}
// 	fmt.Println("write ", n, "bytes in file!!!")
// 	fmt.Println("File append!!")
// }

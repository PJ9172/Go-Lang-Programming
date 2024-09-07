//Write a program to check if a given string is a palindrome (reads the same forwards and backwards).

package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	fmt.Print("Enter the string : ")
// 	reader := bufio.NewReader(os.Stdin)
// 	str, _ := reader.ReadString('\n')
// 	fmt.Println("Checking given string is palindrom :", isPalindrom(str))
// }

// func isPalindrom(str string) bool {
// 	str = strings.ReplaceAll(str, "\n", "")
// 	char := []rune(str)
// 	for i, j := 0, len(char)-1; i < j; i, j = i+1, j-1 {
// 		if char[i] != char[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

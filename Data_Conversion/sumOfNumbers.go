//Write a program that takes a string of comma-separated numbers, converts them to integers, and returns their sum.

package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Enter the string of comma-separated numbers : ")
// 	reader := bufio.NewReader(os.Stdin)
// 	str, _ := reader.ReadString('\n')
// 	str = strings.ReplaceAll(str, ",", "")
// 	str = strings.TrimSpace(str)
// 	num, _ := strconv.Atoi(str)
// 	sum := 0
// 	var r int
// 	for num != 0 {
// 		r = num % 10
// 		sum += r
// 		num /= 10
// 	}
// 	fmt.Println("Sum of numbers is :", sum)
// }

//Write a program that reverses the elements of a slice.

package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print("Enter the size of slice : ")
// 	var size, i, j, t, num int
// 	fmt.Scan(&size)
// 	var slice = []int{}
// 	fmt.Print("Enter the values in slice : ")
// 	for i = 0; i < size; i++ {
// 		fmt.Scan(&num)
// 		slice = append(slice, num)
// 	}
// 	fmt.Println("Slice before reversing  :", slice)
// 	for i, j = 0, size-1; i < j; i, j = i+1, j-1 {
// 		t = slice[i]
// 		slice[i] = slice[j]
// 		slice[j] = t
// 	}
// 	fmt.Print("Slice after reversing : ", slice)
// }

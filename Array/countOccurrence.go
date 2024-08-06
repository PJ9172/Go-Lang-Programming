//Write a program that counts the occurrences of a specific element in an array.

package main

import "fmt"

func main() {
	fmt.Print("Enter the 5 numbers in array : ")
	var arr [5]int
	var i int
	var count = 0
	for i = 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Print("Enter the number to find occurrence : ")
	var num int
	fmt.Scan(&num)
	for i = 0; i < 5; i++ {
		if arr[i] == num{
			count++
		}
	}
	fmt.Print("Occurrence of ",num," is : ",count)
}

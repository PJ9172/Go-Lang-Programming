//Write a program that defines and uses an anonymous function to calculate the square of a number.

package main

import "fmt"

func main() {
	fmt.Print("Enter the number : ")
	var num int
	fmt.Scan(&num)
	square := func(num int) int {
		return num * num
	}
	fmt.Print("Square of given number is : ", square(num))
}

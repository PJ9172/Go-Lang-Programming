//Write a program that converts a date string in the format YYYY-MM-DD to a Go time.Time object and vice versa.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Print("Enter the date in string with format YYYY-MM-DD : ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in Reading", err)
		return
	}
	str = strings.TrimSpace(str)
	layout := "2006-01-02"
	date, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error in date conversion", err)
		return
	}
	fmt.Println("Date is :", date)
	fmt.Printf("Type of Date is %T\n", date)
}

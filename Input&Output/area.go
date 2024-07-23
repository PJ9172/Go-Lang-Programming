// Area of circle,Square & Rectangle

package main

import "fmt"

func main() {

	var r float64
	fmt.Print("Enter the redius of circle : ")
	fmt.Scan(&r)
	fmt.Println("Area of circle :", (3.14 * r * r))

	var s int
	fmt.Print("Enter the side of Square : ")
	fmt.Scan(&s)
	fmt.Printf("Area of Square : %d\n", s*s)

	var l, b int
	fmt.Print("Enter the lenght & breadth : ")
	fmt.Scan(&l, &b)
	fmt.Println("Area of Rectangle :", l*b)
}

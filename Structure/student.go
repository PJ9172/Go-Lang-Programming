//Create a Student struct with fields for Name, Grades (slice of ints), and Age.
//Write methods to calculate the average grade and determine if the student has passed.

package main

// import "fmt"

// type Student struct {
// 	Name   string
// 	Grades []int
// 	Age    int
// }

// func main() {
// 	var s Student
// 	fmt.Print("Enter the student Name : ")
// 	fmt.Scan(&s.Name)
// 	fmt.Print("Enter the number of subjects : ")
// 	var size int
// 	fmt.Scan(&size)
// 	fmt.Print("Enter the grades of subjects : ")
// 	var m int
// 	for i := 0; i < size; i++ {
// 		fmt.Scan(&m)
// 		s.Grades = append(s.Grades, m)
// 	}
// 	fmt.Print("Enter the student Age : ")
// 	fmt.Scan(&s.Age)
// 	sum := 0
// 	for _, value := range s.Grades {
// 		sum += value
// 	}
// 	avg := sum / size
// 	if avg >= 35 {
// 		fmt.Print("Student is passed!!!")
// 	} else {
// 		fmt.Print("Student is failed!!!")
// 	}
// }

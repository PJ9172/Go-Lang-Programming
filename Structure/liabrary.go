//Create a Book struct with fields for Title, Author, Pages, and ISBN.
//Write a program to manage a collection of books, including adding, removing, and searching for books by title.

package main

// import "fmt"

// type Liabrary struct {
// 	Title  string
// 	Author string
// 	Pages  int
// }

// func main() {
// 	var l1 = []Liabrary{
// 		{"abc", "abc", 60},
// 		{"pqr", "pqr", 100},
// 		{"lmn", "lmn", 150},
// 	}
// 	var ch int
// 	fmt.Println("Enter you choice for liabrary management : ")
// 	fmt.Println("1 for adding new book!!")
// 	fmt.Println("2 for removing book!!")
// 	fmt.Println("3 for searching a book!!!")
// 	fmt.Scan(&ch)
// 	switch ch {
// 	case 1:
// 		fmt.Print("Enter the book name : ")
// 		var name, author string
// 		var page int
// 		fmt.Scan(&name)
// 		t := 0
// 		for _, value := range l1 {
// 			if value.Title == name {
// 				t = 1
// 			}
// 		}
// 		if t == 1 {
// 			fmt.Print("Book already exist!!!")
// 		} else {
// 			fmt.Print("Enter the book author : ")
// 			fmt.Scan(&author)
// 			fmt.Print("Enter the number of pages : ")
// 			fmt.Scan(&page)
// 			var l Liabrary
// 			l.Title = name
// 			l.Author = author
// 			l.Pages = page
// 			l1 = append(l1, l)
// 			fmt.Println("Book added successfully!!!")
// 			fmt.Printf("Now Book records are : %v\n", l1)
// 		}
// 	case 2:
// 		fmt.Print("Enter the name of book to remove : ")
// 		var name string
// 		fmt.Scan(&name)
// 		t := 0
// 		for index, value := range l1 {
// 			if value.Title == name {
// 				l1 = append(l1[:index], l1[index+1:]...)
// 				t = 1
// 				break
// 			}
// 		}
// 		if t == 0 {
// 			fmt.Print(name, " book is doesn't exist!!!")
// 		} else {
// 			fmt.Println(name, "book is successfully removed & remaining books are :", l1)
// 		}
// 	case 3:
// 		fmt.Print("Enter the book name to search : ")
// 		var name string
// 		fmt.Scan(&name)
// 		t := 0
// 		for _, value := range l1 {
// 			if value.Title == name {
// 				fmt.Print(name, " book is founded & the details are : ", value)
// 				t = 1
// 			}
// 		}
// 		if t == 0 {
// 			fmt.Print(name, " book not found!!")
// 		}
// 	default:
// 		fmt.Println("Invalid choice!!!")
// 		fmt.Print("Try again!!!")
// 	}
// }

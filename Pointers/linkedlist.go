//Implement a simple linked list in Go using pointers.
//Include functions to add nodes, remove nodes, and print the list.

package main

// import (
// 	"fmt"
// )

// type Node struct {
// 	Data int
// 	Next *Node
// }

// func main() {
// 	var head *Node
// 	var c = 1
// 	var ch int
// 	fmt.Println("Choose the following : ")
// 	fmt.Println("1 for Add nodes!!!")
// 	fmt.Println("2 for Remove nodes!!!")
// 	fmt.Println("3 for Display the list!!!")
// 	fmt.Println()
// 	for c == 1 {
// 		fmt.Print("Enter you choice : ")
// 		fmt.Scan(&ch)
// 		switch ch {
// 		case 1:
// 			head = add(head)
// 		case 2:
// 			head = rem(head)
// 		case 3:
// 			display(head)
// 		default:
// 			fmt.Println("Invalid choice!!!!")
// 		}
// 		fmt.Print("\nDo you want to perform more operations (1/0) : ")
// 		fmt.Scan(&c)
// 	}
// }
// func add(head *Node) *Node {
// 	var temp *Node
// 	var ch = 1
// 	for ch == 1 {
// 		newNode := new(Node)
// 		fmt.Print("Enter the data : ")
// 		fmt.Scan(&newNode.Data)
// 		newNode.Next = nil
// 		if head == nil {
// 			head = newNode
// 			temp = newNode
// 		} else {
// 			temp.Next = newNode
// 			temp = newNode
// 		}
// 		fmt.Print("Do you want to add node (1/0) : ")
// 		fmt.Scan(&ch)
// 	}
// 	return head
// }

// func rem(head *Node) *Node {
// 	fmt.Print("Enter the position of node to delete : ")
// 	var p int
// 	fmt.Scan(&p)
// 	temp := head
// 	if p == 1 {
// 		head = temp.Next
// 	} else {
// 		for i := 1; i < p-1; i++ {
// 			temp = temp.Next
// 		}
// 		t := temp
// 		temp = temp.Next
// 		t.Next = temp.Next
// 	}
// 	fmt.Println("Node removed successfully!!")
// 	return head
// }

// func display(head *Node) {
// 	var temp *Node
// 	for temp = head; temp.Next != nil; temp = temp.Next {
// 		fmt.Print(temp.Data, "->")
// 	}
// 	fmt.Print(temp.Data)
// }

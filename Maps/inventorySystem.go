//Create a simple inventory management system using a map
//where the key is the item name, and the value is the quantity.

package main

import "fmt"

func main() {
	fmt.Print("Enter the number of items in map : ")
	var size, i, quantity int
	fmt.Scan(&size)
	var itemname string
	inventory := make(map[string]int)
	fmt.Print("Enter the item name as key & quantity as value : ")
	for i = 0; i < size; i++ {
		fmt.Scan(&itemname)
		fmt.Scan(&quantity)
		inventory[itemname] = quantity
	}
	fmt.Print("Inventory is : ",inventory)
}

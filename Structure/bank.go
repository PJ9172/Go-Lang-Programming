//Create a BankAccount struct with fields for AccountHolderName, AccountNumber, and Balance.
//Write methods to deposit, withdraw, and check the balance.

package main

// import "fmt"

// type BankAccount struct {
// 	Aname   string
// 	Ano     int
// 	Balance float64
// }

// func main() {
// 	fmt.Print("Enter the details of account holder !!")
// 	var b1 BankAccount
// 	b1.Aname = "Rohit"
// 	b1.Ano = 1234567890
// 	b1.Balance = 50000.50
// 	fmt.Println("Please select following operations : ")
// 	fmt.Println("1 for deposit!!!")
// 	fmt.Println("2 for withdraw!!!")
// 	fmt.Println("3 for check balance!!!")
// 	var ch int
// 	fmt.Scan(&ch)
// 	switch ch {
// 	case 1:
// 		deposit(&b1)
// 	case 2:
// 		withdraw(&b1)
// 	case 3:
// 		balance(&b1)
// 	default:
// 		fmt.Print("Invalid choice!! Try again!!!")
// 	}
// }
// func deposit(b *BankAccount) {
// 	fmt.Print("Enter the amount to deposit : ")
// 	var amt float64
// 	fmt.Scan(&amt)
// 	b.Balance += amt
// 	fmt.Println("Amount deposit successfully!!!")
// 	fmt.Print("Account holder info : ", *b)
// }
// func withdraw(b *BankAccount) {
// 	fmt.Print("Enter the amount to withdraw : ")
// 	var amt float64
// 	fmt.Scan(&amt)
// 	if amt > b.Balance {
// 		fmt.Print("Insuficient balace!!!")
// 	} else {
// 		b.Balance -= amt
// 		fmt.Println("Amount withdraw successfully!!!")
// 		fmt.Print("Account holder info : ", *b)
// 	}
// }
// func balance(b *BankAccount) {
// 	fmt.Print("Your bank balace is : ", b.Balance)
// }

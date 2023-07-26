package main

import "fmt"

func main() {
	// get 2 different bank account balances
	var balance1 float32
	var balance2 float32

	fmt.Println("Input the first bank account's balance:")
	fmt.Scanln(&balance1)

	fmt.Println("Input the second bank account's balance:")
	fmt.Scanln(&balance2)

	// add them together
	newBalance := balance1 + balance2

	// print out the result
	fmt.Printf("The two balances added together is: %.2f\n", newBalance)
}

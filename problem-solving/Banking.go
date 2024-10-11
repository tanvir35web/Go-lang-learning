// Problem description:
// Create a program in Go that models a simple bank account.
// The account should have methods for depositing, withdrawing, and checking the balance.

package main

import (
	"errors"
	"fmt"
)

// BankAccount struct
type BankAccount struct {
	owner   string
	balance float64
}

// method to deposit money to bank account
func (b *BankAccount) Diposit(amount float64) {
	if amount > 0 {
		b.balance += amount
		fmt.Printf("Diposited : $ %.2f \n", amount)
	} else {
		fmt.Println("Deposit amount must be greater than Zero!")
	}
}

// method to withdraw money from bank account
func (b *BankAccount) withdraw(amount float64) error {
	if amount > b.balance {
		return errors.New("Insufficient balance! \n")
	}
	b.balance -= amount
	fmt.Printf("Withdrew : $ %.2f \n", amount)
	return nil
}

// method to check the current balance
func (b *BankAccount) CheckBalance() {
	fmt.Printf("Current Balance: $ %.2f \n", b.balance)
}

func main() {
	//create a bank account instance
	account := BankAccount{
		owner:   "Tanvir",
		balance: 50000.00,
	}

	// Firstly check balance
	account.CheckBalance()

	//deposit money
	account.Diposit(25000)

	//check balance
	account.CheckBalance()

	//withdraw money
	account.withdraw(15000)

	//check balance
	account.CheckBalance()

	//withdraw more money than available
	err := account.withdraw(80000)
	if err != nil {
		fmt.Println(err)
	}
	//check balance
	account.CheckBalance()

}

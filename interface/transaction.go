package main

import "fmt"

// account holder

type User struct {
	name    string
	balance float32
	history []string
}

// payment interface

type Payer interface {
	pay(amount float32, name string) bool
}
type Refunder interface {
	refund(amount float32)
}

type UPI struct{}
type Stripe struct{}

type PaymentServices struct {
	method Payer
}

type RefundServices struct {
	method Refunder
}

func (u *User) Deposite(amount float32) {
	if amount <= 0 {
		fmt.Println("you can deposite 0 or negative money")
		return 
	}
	u.balance += amount
	u.history = append(u.history, fmt.Sprintf("deposited: %.2f", amount))

}
func (u User) showHistory() {

	fmt.Println("\ntransaction History")
	for _, val := range u.history {
		fmt.Println("-", val)
	}
}

func main() {

	user := User{name: "amit", balance: 0}
    
	user.Deposite(500)
	user.Deposite(5070)
	user.Deposite(5600.70)
	user.showHistory()

}

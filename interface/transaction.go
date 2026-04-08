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

func(s Stripe) pay(amount float32,name string) bool{
      if amount>10000 {
		fmt.Println("you cant send more than 10000")
		return false
	  }
	  fmt.Println("payment done via stripe:",amount)
	  return true
}

func (s Stripe) refund(amount float32){
	fmt.Println("money refund: ",amount)
}

func(p PaymentServices) sendMoney(amount float32,user *User){
	 
	if user.balance < amount {
             fmt.Println("insuficient balance")
			 return
	}
	isSuccess:=p.method.pay(amount,user.name)

	if isSuccess {
		user.balance-=amount
		user.history=append(user.history,fmt.Sprintf("paid : %.2f",amount))
	}else{
		user.balance-=amount
		fmt.Println("transaction failed money will be refund ....")
	}

}



func(r RefundServices) refundMoney(amount float32,user *User){

	user.balance+=amount
	user.history=append(user.history,fmt.Sprintf("ammount refund : %.2f",amount))
	r.method.refund(amount)
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
    
	user.Deposite(50000)
	user.Deposite(5070)
	user.Deposite(5600.70)
	fmt.Println("........................")
	fmt.Println("total balance:",user.balance)
	user.showHistory()
	newStripePayment:=Stripe{}

	p1:=PaymentServices{method:newStripePayment}
	p1.sendMoney(5000,&user)
	p1.sendMoney(20000,&user)
   fmt.Println("........................")
	fmt.Println("total balance:",user.balance)

	fmt.Println("........................")
	user.showHistory()

	fmt.Println("........................")
	fmt.Println("total balance:",user.balance)

	newRefund:=RefundServices{method:newStripePayment}
	newRefund.refundMoney(20000,&user)
	fmt.Println("........................")
	user.showHistory()

	fmt.Println("........................")
	fmt.Println("total balance:",user.balance)

	



}

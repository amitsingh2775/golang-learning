package main

import "fmt"

// important flow 
// Interface → rule banata hai
// Struct → rule follow karta hai
// Method → struct ke andar likhte hain



// it will be make the problem if i need to add more or another payment gatway
type Payment struct{}
type Stripe struct{}

func (s Stripe) pay(amount float32,name string){
	fmt.Println("payment succesfull :-",amount,name)
}

// method

func (p Payment) makePayment(amount float32,name string){
	stripeMethod:=Stripe{}
	stripeMethod.pay(amount,name)
}

func main(){
    newpayment:=Payment{}
	newpayment.makePayment(50000,"amit")
}
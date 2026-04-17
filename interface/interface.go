// package main

// import "fmt"

// type paymentMethod interface{
// 	pay(amunt float32,name string)
// 	refund(amount float32,account_number int)
// }

// type Stripe struct{}
// type UPI struct{}

// func (s Stripe) pay(amount float32 ,name string){
// 	fmt.Println("payment successfull via Stripe :","ammout:",amount,"sender",name)
// }

// func (u UPI) pay(amount float32 ,name string){
// 	  fmt.Println("payment successfull via UPI :","ammout:",amount,"sender",name)
// }

// func (s Stripe) refund(amount float32,account_number int){
// 	fmt.Println("payment refund of account number :",account_number,"amount: ",amount)
// }

// type Payment struct{
// 	method paymentMethod
// }

// func (p Payment) makePayment(amount float32,name string){
// 	p.method.pay(amount,name)
// }

// func (p Payment) refoundMoney(amount float32,account_number int){
// 	p.method.refund(amount,account_number)
// }


// func main(){
//   newstripe:=Stripe{}
// //   newUPI:=UPI{}

//   p1:=Payment{method:newstripe}
//   p1.makePayment(25.66,"amit")
//   p1.refoundMoney(45.6,1233324223)
// //   p2:=Payment{method:newUPI}
// //   p2.makePayment(700.67,"ashish")
// }
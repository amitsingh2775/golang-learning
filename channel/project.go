package main

import  ( "fmt"
          "time")

type Order struct{
	id int
	ammount int
}

 func OrderProcess(OrderChan chan Order){
	for order:=range OrderChan{
		fmt.Println("order processing id is ",order.id)
		time.Sleep(time.Second)
		fmt.Println("Processing Order ",order.id,"ammounr",order.ammount)
	}
 }

func main(){
 
	// creating the order chanal
	OrderChan:=make(chan Order)
	go OrderProcess(OrderChan)

	for i:=1;i<=5;i++{
		order:=Order{id:i,ammount:i*100}
		fmt.Println("order created ",order)
		OrderChan <- order // here sending the order
	}
	close(OrderChan)

	time.Sleep(time.Second*2)

}
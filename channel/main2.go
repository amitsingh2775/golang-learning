package main

import "fmt"


func Sum(result chan int,num1 int,num2 int){
	res:=num1+num2
	result <-res
}

func main(){

	result:=make(chan int)
	go Sum(result,5,3)
	res:=<-result
	fmt.Println(res)
}
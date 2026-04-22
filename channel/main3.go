package main


import "fmt"


func main(){

	Chan1:=make(chan int)
	Chan2:=make(chan string)

	// inline function

	go func(){
		Chan1<- 10
	}()

	go func(){
		Chan2 <- "amit"
	}()


	for i:=0;i<2;i++{
       select{
		  
	   case chan1_val:=<-Chan1: 
	           fmt.Println("chan1 value",chan1_val)
	   case chan2_val:=<-Chan2:
		       fmt.Println("chan2 value ",chan2_val)
	   }
	}


}
 package main

import ("fmt"
        "time"
	     "math/rand")

// it is a data pipe b/w two go routines

// Note:
//  messageChannel <- "hii"   data behjna syntax
// message:=<-messageChannel  data recive karna syntax \

func Recived(numChan chan int){
   for num :=range numChan{
	 fmt.Println(num)
	 time.Sleep(time.Second)
   }
}

func main (){

	numChan:=make(chan int)

	go Recived(numChan)

	for{
		numChan <- rand.Intn(100)
	}

} 
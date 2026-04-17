package main

import ("fmt"
          "time")

//Goroutine = lightweight thread -parallel execution karta hai

//Goroutines = concurrency, not always parallelism
// Concurrency = multiple tasks manage karna
// Parallelism = same time me run hona (depends on CPU)

func printId(id int){
	fmt.Println("user ID,",id)
}

func main(){
   
	for i:=0;i<=9;i++{
		// must to use go keyword to make goroutine
		go printId(i)
	}

	time.Sleep(time.Second*2)
}
package main

import ("fmt"
          "sync")

//Goroutine = lightweight thread -parallel execution karta hai

//Goroutines = concurrency, not always parallelism
// Concurrency = multiple tasks manage karna
// Parallelism = same time me run hona (depends on CPU)

func printId(id int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("user ID,",id)
}

func main(){
    var wg sync.WaitGroup
	for i:=0;i<=9;i++{
		wg.Add(1)
		// must to use go keyword to make goroutine
		go printId(i,&wg)
	}

	wg.Wait()
}
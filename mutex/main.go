//Mutex = lock system
//Ek time pe sirf 1 goroutine ko shared data access karne deta hai and ye race condition ko avoide karta hai


// note defer function hamesha function ke last me call hota hai 
package main

import ("fmt"
         "sync" )


type Post struct{
	like int
	mu sync.Mutex
}

func (p * Post) likeIncr(wg* sync.WaitGroup){
	defer func(){
        wg.Done()
		p.mu.Unlock()
	}() 

	p.mu.Lock()
	p.like+=1  // best practice unlock ko defer function ke sath rakho agar 
	
}
func main(){
   var wg sync.WaitGroup
    post1:=Post{like:0}

	 for i:=0;i<100;i++{
		wg.Add(1)
	   go post1.likeIncr(&wg)
	
	 }
	 wg.Wait()
	 fmt.Println("totla likes: ",post1.like)

	


}
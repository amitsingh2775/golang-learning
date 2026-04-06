package main

import "fmt"

func loginLimiter(maxAttempt int)func() bool{
    var attempt int=0
	
	return func() bool{
		if attempt>=maxAttempt {
			fmt.Println("you have reached max attempt,block for 24hr")
			return false
		}
		attempt+=1
		fmt.Println("attempt: ",attempt)
		return true
	}

}

func main(){
    login:=loginLimiter(3)
	login()
	login()
	login()
	login()
}
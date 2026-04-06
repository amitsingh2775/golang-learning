
package main 

import "fmt"

//Closure = function jo apne outer variables ko yaad rakhta hai (even after outer function khatam ho jaye)

func Counter()func() int{
	var cnt int=0
	return func()int{
		cnt+=1
		return cnt
	}
}

func main (){
    callFunc:=Counter()
	fmt.Println(callFunc())
	fmt.Println(callFunc())
	fmt.Println(callFunc())
	

}
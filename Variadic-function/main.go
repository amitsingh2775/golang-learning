package main

import "fmt"


// Variadic function
// ... ye internalu slice ki trah kam karta hai
func sum(nums ...int)int{
	sum:=0
	for _,val:=range nums{
          sum+=val
	}
	return sum
}
//Multiple parameters + variadic
func data(name string,nums ...int){
	fmt.Println(name,nums)
}

func main(){
  fmt.Println(sum(1,3,5,3,3))
  fmt.Println(sum(1,3,4,6,4,5,3,3))
  data("amit",23,3,3,2,2)
}
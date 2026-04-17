package main


import "fmt"

func printSlice[T any](nums []T){
	for _,data:=range nums{
		fmt.Println(data)
	}
}

 // generics with struct

 type stack [T any]struct{
	ele []T
 }
func main(){
    myStack:=stack[int]{
		ele:[]int{1,2,3},
	}
	fmt.Println(myStack)
	fmt.Println("//////////..............////////////")
	printSlice([]int{1,2,3,4,4})
	fmt.Println("....................................")
	printSlice([]string{"amit","ashish","abhay"})


}
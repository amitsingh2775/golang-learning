package main


import "fmt"

func change(val *int){
	*val+=1
}

func main(){

	var val int=6
    fmt.Println("before change ",val)
	change(&val)
	fmt.Println("after change ",val)
	
}
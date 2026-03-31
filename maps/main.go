package main

import "fmt"

func main(){

	// map declearation

	dataMap:=make(map[int]string)

	// insert value in map
	dataMap[1]="amit"
	dataMap[2]="rahul"

	// get values 

	for _,data:=range dataMap{
		fmt.Println(data)
	}

	// key existing check here its asked most of the interview

	val,ok:=dataMap[2]

	if ok {
		fmt.Println("found and value: ",val)
	}else{
		fmt.Println("not found")
	}
   // deletion in map

   delete(dataMap,2)
   fmt.Println("/.................../")
   for _,data:=range dataMap{
		fmt.Println(data)
	}

}
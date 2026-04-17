package main

import "fmt"


// for integer

type Order int

const (
	Ricived Order=200
	Faild=404
	processing=101
)

func Status(data Order){
	fmt.Println(data)
}

// string type
type Code string

const (
	Username Code="amit"
	Address="sdfsd"
)

func Details(data Code){
	fmt.Println(data)
}
func main(){
//    Status(Ricived)
//    Status(processing)
 var input  string
  fmt.Scanln(&input)
  code:=Code(input)
   Details(code)
}
package main


import "fmt"

func main(){

	var rand int

	fmt.Println("enter a random number between 1 and 5")
	fmt.Scanln(&rand)

	switch rand{
	case 1:

		fmt.Println("you entered 1")
	case 2:
		fmt.Println("you entered 2")
	case 3:		
		fmt.Println("you entered 3")
	case 4:
		fmt.Println("you entered 4")
	case 5:
		fmt.Println("you entered 5")
	default:
		fmt.Println("you entered an invalid number")
	}
}
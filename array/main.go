package main

import "fmt"

func main() {

	// var arr [5]int

	// arr[3] = 30
	// fmt.Println(arr)

	// var arr [10]int

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Scanln(&arr[i])
	// }
	// fmt.Println(arr)

	// using range with index

	// arr := [3]int{3, 5, 3}

	// for index,value:=range arr {
	// 	fmt.Println(index,value)
	// }

	// whitout index
	// for _, value := range arr {
	// 	fmt.Println(value)
	// }

	// size auto infer
	// nums := [...]int{3, 5, 3, 2, 4, 3, 5}
	// fmt.Println(nums)

	//2d array
	// var arr [2][4]int
	// arr[0][0] = 5
	// fmt.Println(arr)

	// nums := [2][2]int{{1, 2}, {2, 4}}
	// fmt.Println(nums)

	var nums [2][2]int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			fmt.Scan(&nums[i][j])
		}

	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			fmt.Print(nums[i][j], " ")
		}
		fmt.Println()
	}

}

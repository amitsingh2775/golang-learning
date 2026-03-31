package main

import "fmt"

func main() {

	// nums := make([]int, 10, 20)
	// for i := 0; i < 10; i++ {
	// 	nums[i] = 2 * i
	// }
	// fmt.Println(nums)

	// var arr[]int
	// arr=append(arr,23,34,3,35,353,35)
	// fmt.Println(arr)

	// shared memeory concept
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	b[0] = 100
	fmt.Println("first array", a)
	fmt.Println("second array", b)
}

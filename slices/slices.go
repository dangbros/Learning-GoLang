package main

import "fmt"

func main() {
	// var arr []int

	var arr = make([]int, 2, 5)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)

	fmt.Println(cap(arr))
	fmt.Println(arr)

}

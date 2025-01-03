package main

import (
	"fmt"
)

func makeForSlice() {
	var intSlice = make([]int, 10)        //here length and capacity are the same
	var strSlice = make([]string, 10, 20) //here lengtha and capacity are different

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
}
func makeForMaps() {
	var employee = make(map[string]int)
	employee["Mark"] = 20
	employee["Rick"] = 10
	fmt.Println(employee)
}

func main() {
	makeForSlice()
	makeForMaps()
}

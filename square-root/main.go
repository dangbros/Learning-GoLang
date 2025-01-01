package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Println("Enter the number: ")
	fmt.Scan(&num)
	fmt.Printf("This is the square root of the number: %v\n", math.Sqrt(num))
}

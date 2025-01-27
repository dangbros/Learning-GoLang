package main

import (
	"fmt"
	"learning_golang/package/average" // Use full module path instead of "./average"
)

func main() {
	var a, b int
	fmt.Println("enter two numbers: ")
	fmt.Scan(&a, &b)
	result := average.Average(a, b)
	fmt.Printf("Average is: %f\n", result)
}

package main

import (
	"fmt"
	"learning_golang/addition/mathutils"
	"learning_golang/addition/mypackage"
	"os"
)

func main() {
	var a, b int
	env := os.Environ()
	fmt.Println("the environment is:", env)
	fmt.Println("Enter the 2 numbers: ")
	fmt.Scan(&a, &b)
	sum := mypackage.Add(a, b)
	product := mathutils.Multi(a, b)
	fmt.Printf("Sum of %v and %v is %v\n", a, b, sum)
	fmt.Printf("Product of %v and %v is %v\n", a, b, product)

}

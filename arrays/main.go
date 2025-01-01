package main

import "fmt"

func matrix() {
	var mat [3][3]int
	mat[0] = [3]int{0, 1, 0}
	mat[1] = [3]int{1, 0, 0}
	mat[2] = [3]int{0, 0, 1}
	fmt.Println(mat)
}

func main() {
	arr := [...]int{87, 91, 78} //... sets the size of the array automatically
	student := [3]string{}
	fmt.Printf("value of array: %v\n", arr)
	fmt.Printf("Students: %v\n", student)
	student[0] = "Lisa"
	fmt.Printf("Student 1: %v\n", student)
	matrix()
}

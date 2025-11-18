package main

import "fmt"

func main() {
	// nums := []int{1, 2, 6, 7, 8}
	// for i := 0; i < len(nums); i++ {
	// fmt.Println(nums[i])
	// }

	// for i, num := range nums {
	// f
	// mt.Println(num, i)
	// }

	// m := map[string]string{"fname": "john ", "lname": "doe"}

	// for k, v := range m {
	// fmt.Println(k, v)
	// }

	// for k := range m {
	// fmt.Println(k)
	// }

	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}

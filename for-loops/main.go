package main

import (
	"fmt"
)

func main() {
	arr := [...]int{87, 91, 78}
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for _, value := range arr {
		fmt.Println(value)
	}

}

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th cleaning day")
	case 12:
		fmt.Println("Today is 12th, visit doctor")
	case 17:
		fmt.Println("Today is 17th, go out to eat dinner")
	case 28:
		fmt.Println("Today is 28th, grab a drink at bar with friends")
	case 31:
		fmt.Println("Today is 31th, Meet your parents")
	default:
		fmt.Println("No information available")
	}
}

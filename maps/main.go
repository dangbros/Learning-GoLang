package main

import "fmt"

func main() {
	var mpp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	fmt.Println(mpp)
	mpp["apple"] = 900        //applies to change any value inside the map
	fmt.Println(mpp["apple"]) //updated apple value
	//deleting a key and value in map
	delete(mpp, "apple") //using the delete function
	fmt.Println(mpp)
	//check if a key exist in the map
	val, ok := mpp["apple"] //val will store the value of the key and ok will store if the key exist or not
	fmt.Println(val, ok)
	val2, ok2 := mpp["pear"]
	fmt.Println(val2, ok2)
}

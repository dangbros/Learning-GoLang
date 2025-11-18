package main

import "fmt"

func main() {
	//creating map
	m := make(map[string]string)

	//setting elements
	m["name"] = "golang"
	m["area"] = "backend"

	//getting elements
	fmt.Println(m["name"])
	fmt.Println(m["area"])

	//if key value doesn't exists it return null
	fmt.Println(m["phone"])

	m2 := make(map[string]int)

	m2["age"] = 20
	fmt.Println(m2["age"])
	fmt.Println("lenght of maps: ", len(m), len(m2))

	//deleting data from map
	fmt.Println("---deleting data from map---")
	fmt.Println("before: ", m)
	delete(m, "area")
	fmt.Println("after: ", m)

	//clearing map
	clear(m2)
	fmt.Println(m2)

	//creating map alternative way
	m3 := map[string]int{"price": 40}
	fmt.Println(m3)
	// var m4 map[string]string
}

package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; No super or parent
	sourya := User{"Sourya", "souryaroy30@gmail.com", true, 19}
	fmt.Println(sourya)
	fmt.Printf("Details: %+v\n", sourya)
	fmt.Printf("Name is %v and email is %v.\n", sourya.Name, sourya.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

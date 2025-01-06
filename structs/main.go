package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string `validate:"min=2,max=32"`
	Email string `validate:"required,email"`
}

func validate(val interface{}) error {
	v := reflect.ValueOf(val)
	fmt.Println(v)
	return nil
}

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; No super or parent
	sourya := User{"Sourya", "souryaroy30@gmail.com", true, 19}
	fmt.Println(sourya)
	fmt.Printf("Details: %+v\n", sourya)
	fmt.Printf("Name is %v and email is %v.\n", sourya.Name, sourya.Email)

	rick := User{
		Name:  "Rick",
		Email: "royshourya30@proton.me",
	}

	fmt.Println(validate(rick))

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

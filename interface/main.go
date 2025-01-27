package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Mover interface {
	Move() string
}

type Robot struct {
	Name         string
	ID           int
	MovementType string
}

type Dog struct {
	Name string
}

func (r Robot) Speak() string {
	return fmt.Sprintf("Hello Bee Bop I am %v, my ID number is %v", r.Name, r.ID)
}

func (r Robot) Move() string {
	return fmt.Sprintf("%v is moving using %v, Beep Bop", r.Name, r.MovementType)
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return "woof!, I am " + d.Name
}

func (c Cat) Speak() string {
	return "mewww!, I am " + c.Name
}

func main() {
	speakers := []Speaker{Dog{Name: "Jimmy"}, Cat{Name: "Kitty"}}
	robots := []Robot{
		{
			Name:         "Rick",
			ID:           2024,
			MovementType: "treads",
		},

		{
			Name:         "Mittu",
			ID:           2025,
			MovementType: "rocket-boosters",
		},
	}
	for _, robo := range robots {
		fmt.Println(robo.Speak())
		fmt.Println(robo.Move())
		fmt.Println("")
	}

	for _, s := range speakers {
		fmt.Println(s.Speak())
	}

}

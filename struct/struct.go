package main

import (
	"fmt"
	"time"
)

// this is how we created the struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// this is the constructor
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder // we return the address of the constructor
}

// this is function of the struct
func (o *order) updateStatus(status string) {
	o.status = status
}

// this is also the function of the struct
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	//this is how we can create a struct and instantiate and assign values
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
	// myOrder := newOrder("1", 50, "received")
	// fmt.Println(myOrder.amount)

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50,
	// 	status: "received",
	// }

	// myOrder.updateStatus("returned")
	// fmt.Println(myOrder)

	// fmt.Println(myOrder.getAmount())

	// myOrder.createdAt = time.Now()

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"
	// fmt.Println("Order1 struct has", myOrder)
	// fmt.Println("Order2 struct has", myOrder2)
}

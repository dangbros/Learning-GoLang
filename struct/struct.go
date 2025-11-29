package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *order) updateStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	myOrder := order{
		id:     "1",
		amount: 50,
		status: "received",
	}

	myOrder.updateStatus("returned")
	fmt.Println(myOrder)

	fmt.Println(myOrder.getAmount())

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

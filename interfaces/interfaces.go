package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
	// razorPayPaymentGW := razorPay{}
	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)
	// razorPayPaymentGW.pay(amount)
}

type razorPay struct{}

func (r razorPay) pay(amount float32) {
	fmt.Printf("payment of %v is done using razorpay.", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Printf("payment of %v is done using stripe.", amount)
// }

type fakePay struct{}

func (f fakePay) pay(amount float32) {
	fmt.Printf("payment of %v is done using fake-pay!", amount)
}

func main() {
	// stripePaymentGw := stripe{}
	// razorPayGW := razorPay{}
	fakePayGw := fakePay{}

	newPayment := payment{
		gateway: fakePayGw,
	}
	newPayment.makePayment(100)
}

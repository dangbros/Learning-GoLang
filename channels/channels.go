package main

import "fmt"

// func processNum(numChan chan int) {
// 	for nums := range numChan {
// 		fmt.Println("processing number", nums)
// 		time.Sleep(time.Second)
// 	}
// }

//	func add(result chan int, num1 int, num2 int) {
//		sum := num1 + num2
//		result <- sum
//	}

// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("processing...")
// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "ping"
	}()

	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := range 10 {
	// 	emailChan <- fmt.Sprintf("%v@example.com", i)
	// }

	// fmt.Println("done sending!")
	// close(emailChan) //this is important!
	// <-done

	// emailChan <- "a@example.com"
	// emailChan <- "b@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	// coffeCounter := make(chan string)

	// go func() {
	// 	fmt.Println("Brewing coffee...")
	// 	time.Sleep(2 * time.Second)

	// 	fmt.Println("Order ready! Sending to counter...")
	// 	coffeCounter <- "Espresso!"
	// }()

	// fmt.Println("Customer waiting...")

	// cup := <-coffeCounter

	// fmt.Println("got my cup of ", cup)

	// done := make(chan bool)
	// go task(done)
	// <-done

	// result := make(chan int)

	// go add(result, 4, 5)
	// sum := <-result //blocking

	// fmt.Println(sum)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.IntN(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "hello"

	// msg := <- messageChan

	// fmt.Println(msg)
}

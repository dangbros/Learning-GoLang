package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // defer is used to execute a code at the end of the function so the done method is called after the execution of task.
	fmt.Println("doing task id -", id)
}

func main() {
	var wg sync.WaitGroup // this is where we intsantitate the wait group for the go routine

	for i := range 10 {
		wg.Add(1)         // here it acts as a counter on how many light weight threads(go-routines) are running.
		go task(i+1, &wg) // we have to pass the address of the wait group
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Hi!", done)
	// dones[1] = make(chan bool)
	go greet("You!", done)
	// dones[2] = make(chan bool)
	go slowGreet("Blah ... blah ... blah", done)
	// dones[3] = make(chan bool)
	go greet("Have a nice day!", done)

	// for _, done := range dones {
	// 	<-done
	// }

	// for doneChan := range done {
	// 	fmt.Println(doneChan)
	// }

	for range done {
		// 	fmt.Println(doneChan)
	}
}

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
}

func main() {
	done := make(chan bool)

	go greet("Hi!", done)
	go greet("You!", done)
	go slowGreet("Blah ... blah ... blah", done)
	go greet("Have a nice day!", done)

	<-done
	<-done
	<-done
	<-done
}

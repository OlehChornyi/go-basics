package main

import "fmt"

type customString string

// Reciever
func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString = "Oleh"

	name.log()
}

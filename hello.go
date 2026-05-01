package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Hello())


	var firstName,lastName string 

	firstName = "Leon"
	lastName = "Jose"

	fmt.Printf("My name is %s %s\n", firstName, lastName)
}



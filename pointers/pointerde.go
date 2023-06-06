package main

import "fmt"

func main() {
	b := 245
	a := &b 

	fmt.Println("Value using pointers", *a)
	fmt.Println("Address using pointers", a)
}
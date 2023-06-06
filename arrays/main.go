package main

import (
	"fmt"
	"array/simpleinterest"
)

func main(){
	fmt.Println("This is from array module")
	p := 5000.0
	r := 10.0
	t := 3.0
	si := simpleinterest.Calculate(p,r,t)
	fmt.Println("Simple interest is", si)
}
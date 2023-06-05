package main

import (
	"fmt"
	"unsafe"
)

func main(){
	a := 89
	b := 96
	fmt.Println("type of a is %T and sise of a is %d", unsafe.Sizeof(a))
	fmt.Println("type of b is %T and size of b is %d", unsafe.Sizeof(b))
}
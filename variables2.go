package main 

import (
	"fmt"
	"math"
)

func main() {
	a,b := 145.1, 146.5
	c := math.Min(a,b)
	fmt.Println("Min",c )
}
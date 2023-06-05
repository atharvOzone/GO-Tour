package main

import "fmt"

func rectProps2(l,w float64) (area, per float64) {
	area = l*w
	per = 2 * (l+w)
	return
}

func main(){
	a,_ := rectProps2(12,10)
	fmt.Println("Area", a, "Perimeter")
}
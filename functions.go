package main

import "fmt"

func rectProps(length, width float64) (float64,float64){
	var area = length * width
	var per = (length+width) * 2
	return area, per
}

func main() {
	area, per := rectProps(12,10)
	fmt.Println("area", area, "perimeter", per)
}
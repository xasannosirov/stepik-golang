package main

import "fmt"

func convert(n int64) uint16 {
	return uint16(n)
}

func main() {
	var number int64 = 23
	numberConv := convert(number)
	fmt.Printf("numberConv = %T\nnumber = %T\n", numberConv, number)
}

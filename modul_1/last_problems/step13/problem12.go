package main

import (
	"fmt"
	"math"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	var n int
	fmt.Scan(&n)
	for temp := 0; n >= powInt(2, temp); temp++ {
		fmt.Print(powInt(2, temp), " ")
	}
}

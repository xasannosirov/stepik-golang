package main

import (
	"fmt"
	"math"
)

func pifagor(a int, b int) (int, error) {
	return int(math.Sqrt(float64(a*a + b*b))), nil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res, _ := pifagor(a, b)
	fmt.Println(res)
}

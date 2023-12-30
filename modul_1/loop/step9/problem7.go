package main

import (
	"fmt"
)

func main(){
	var x, p, y, day int
	fmt.Scan(&x, &p, &y)

	for x < y {
		x += (x * p) / 100
		day += 1
	}
	fmt.Println(day)
}
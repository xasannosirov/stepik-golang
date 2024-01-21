package main

import (
	"fmt"
)

func main() {
	var n, h, m int
	fmt.Scan(&n)
	h = n / 3600
	n = n - (h * 3600)
	m = n / 60
	fmt.Println("It is", h, "hours", m, "minutes.")
}

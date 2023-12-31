package main

import (
	"fmt"
)

func main() {
	var n, res int
	fmt.Scan(&n)
	for n%10 > 0 {
		res = res*10 + (n % 10)
		n /= 10
	}
	fmt.Println(res)
}

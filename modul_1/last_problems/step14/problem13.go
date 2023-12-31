package main

import (
	"fmt"
)

func fibonacci(n int) int {
	var f1, f2, temp int = 1, 1, 1
	for f1 <= n {
		if f1 == n {
			return temp
		}
		temp += 1
		f1, f2 = f2, f1+f2
	}
	return -1
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

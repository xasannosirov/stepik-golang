package main

import "fmt"

func iteration(a int) int {
	res := 0
	for a%10 > 0 {
		res += a % 10
		a /= 10
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	for n > 10 {
		n = iteration(n)
	}
	fmt.Println(n)
}

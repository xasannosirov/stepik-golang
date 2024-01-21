package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res := a - 1
	for i := a; i <= b; i++ {
		if i%7 == 0 && b >= i && i >= a {
			res = i
		}
	}
	if res < a {
		fmt.Println("NO")
	} else {
		fmt.Println(res)
	}
}

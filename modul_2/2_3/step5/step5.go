package main

import "fmt"

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

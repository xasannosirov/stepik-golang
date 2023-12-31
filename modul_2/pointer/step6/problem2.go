package main

import "fmt"

// ikkita o'zgaruvchini qiymatini lamashtirib qo'yadi
func test(x1 *int, x2 *int) {
	tempx := *x1
	tempy := *x2
	*x1 = tempy
	*x2 = tempx
	fmt.Println(*x1, *x2)
}

func main() {
	num1, num2 := 9, 2
	test(&num1, &num2)
}

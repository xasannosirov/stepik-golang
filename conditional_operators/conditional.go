package main

import "fmt"

func conditionalOperators() {
	a, b := 8, 3
	num1 := a > b //kattami?
	num2 := a < b //kichikmi?
	num3 := a >= b //katta yoki teng
	num4 := a <= b //kichik yoki teng
	num5 := a == b //teng
	num6 := a != b //teng emas
	fmt.Println(num1, num2, num3, num4, num5, num6) //true false true false false true
}

func main() {
	conditionalOperators()
}

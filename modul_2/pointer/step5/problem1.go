package main

import "fmt"

// ikkita o'zgaruvchini pointeri orqali 
// qiymatiga murojat qilib yodamida ko'paytiradi
func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

func main(){
	num1, num2 := 9, 2
	test(&num1, &num2)
}
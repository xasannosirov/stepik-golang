package main

import "fmt"

func main() {
	var n, summa int
	fmt.Scan(&n)
	for n%10 > 0 {
		summa += n % 10
		n /= 10
	}
	fmt.Println(summa)
}

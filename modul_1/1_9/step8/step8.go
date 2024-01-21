package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	var n1, n2 int
	n1 = (number % 10) + ((number % 100) / 10) + ((number / 100) % 10)
	n2 = ((number / 1000) % 10) + ((number / 10000) % 10) + ((number / 100000) % 10)

	if n1 == n2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

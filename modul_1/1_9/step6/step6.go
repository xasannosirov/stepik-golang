package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	var n1, n2, n3 int = number / 100, (number % 100) / 10, number % 10

	if n1 != n2 && n1 != n3 && n2 != n3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

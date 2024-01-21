package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println("Число положительное")
	} else if number < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}

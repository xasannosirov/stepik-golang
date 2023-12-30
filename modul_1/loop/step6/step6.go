package main

import "fmt"

func BreakOperator() {
	var sum = 0

	for i := 1; i <= 9; i++ {
		if i > 4 {
			break // i 4 dan katta bo'lganida tsikl to'xtaydi
		}
		sum += i
	}
	fmt.Println("Сумма: ", sum) // 10
}

func ContinueOperator() {
	var sum = 0

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // agar i juft bo'lsa yana tsikl shartiga qaytadi
		}
		sum += i
	}
	fmt.Println("Сумма: ", sum) // 25
}

func main() {
	BreakOperator()
	ContinueOperator()
}

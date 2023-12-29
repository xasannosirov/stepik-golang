package main

import "fmt"

func main() {
	ThisFor()
}

func ThisFor() {
	// The sum of numbers from 1 to 10
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) //45

	// Square the numbers from 1 to 10
	var i = 1
	for i < 10 {
		fmt.Println(i * i)
		i++
	}

	var n int
	// It continues until you enter 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)
	}

	//contunue operator
	sum = 0
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // move on to the next iteration
		}
		sum += i
	}
	fmt.Println("Сумма: ", sum) // sum: 25

	// break operator
	sum = 0
	for i := 1; i <= 9; i++ {
		if i > 4 {
			break // if the number is greater than 4 we exit the loop
		}
		sum += i
	}
	fmt.Println("Сумма: ", sum) // sum: 10
}

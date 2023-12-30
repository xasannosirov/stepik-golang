package main

import "fmt"

func main() {
	var number int
    fmt.Scan(&number)

	for number / 10 > 0{
		number /= 10
	}
	fmt.Println(number%10)
}
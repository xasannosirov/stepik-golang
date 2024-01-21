package main

import "fmt"

func main() {
	var a, b int
	array := []int{}
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		array = append(array, b)
	}
	for i := 0; i < a; i++ {
		if i%2 == 0 {
			fmt.Print(array[i], " ")
		}
	}
}

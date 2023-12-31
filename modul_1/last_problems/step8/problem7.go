package main

import "fmt"

func main() {
	var a, b, summa int
	array := []int{}
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		array = append(array, b)
	}
	for i := 0; i < a; i++ {
		if array[i] == 0 {
			summa += 1
		}
	}
	fmt.Println(summa)
}

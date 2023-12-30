package main

import "fmt"

func main() {
	// arrayga qiymatlar qo'shish
	a := []int{1, 2, 3}
	a = append(a, 4, 5)
	fmt.Println(a) // [1 2 3 4 5]
}

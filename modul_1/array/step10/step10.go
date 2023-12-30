package main

import "fmt"

// elementlarni o'chirish
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a) // [1 2 4 5 6 7]
}

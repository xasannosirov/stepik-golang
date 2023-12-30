package main

import "fmt"

// array  uzunligi berilsa qiymati o'garmaydi
func fnA(a [3]int) {
	a[1] = 15
}

// array uzunligi berilmasa qiymati o'zgaradi
func fnB(a []int) {
	a[1] = 15
}

func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}

	fnA(a)
	fnB(b)

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 15 3]
}

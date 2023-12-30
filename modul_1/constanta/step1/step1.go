package main

import (
	"fmt"
)

// keyingi o'zgaruvchiga qiymat berilmasa oldingisi olishi
const (
	A int = 45
	B
	C float32 = 3.3
	D
)

// constantani e'lon qilish
const (
	a int = 45
	b float32 = 3.3
)

func main() {
	fmt.Println(A, B, C, D) // Вывод: 45 45 3.3 3.3
	fmt.Println(a,b)
}

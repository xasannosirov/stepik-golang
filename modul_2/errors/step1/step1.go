package main

import "fmt"

func divide(a int, b int) int {
	return a / b
}

// fmt.Scan() funksiyasi o'qib olishda xatolik yuz berishini tekshirish
func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Ma'lumotni o'qib olishda xatolik yuz berdi")
	} else {
		fmt.Println(divide(input, 5))
	}
}

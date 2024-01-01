package main

import "fmt"

func ShowElementInMap() {
	m := map[int]int{
		12: 2,
		1:  5,
	}
	// mapdagi qiymatni ko'rish
	fmt.Println(m[12]) // 2
}

func DeleteInMap() {
	n := map[int]int{
		12: 2,
		1:  5,
	}
	// mapdan ma'lumotlarni o'chirish
	delete(n, 12)  // 12 kaliti yodamida mapdan ma'lumotni o'chirib yuborish
	fmt.Println(n) // map[1:5]
}

func CheckInMap() {
	m := map[int]int{
		1: 10,
	}
	// mapda ma'lum bir qiymat borligini tekshirish
	if value, inMap := m[1]; inMap {
		fmt.Println(value) // 10
	}

	if value, inMap := m[2]; inMap {
		fmt.Println(value) // ifga kirmasa hechnarsa chiqmaydi
	}
}

func ShowAllElement() {
	m := map[int]int{
		12: 2,
		1:  5,
	}
	// mapdagi barcha ma'lumotlarni ko'rish
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func ShowMapLen() {
	// mapda nechta element borligini tekshirish
	m := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}
	fmt.Println(len(m)) // 3
}

func main() {
	ShowElementInMap()

	DeleteInMap()

	CheckInMap()

	ShowAllElement()

	ShowMapLen()
}

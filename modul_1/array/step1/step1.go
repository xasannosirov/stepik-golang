package main

import "fmt"

func arrayDeclered() {
	var a [3]int //array e'lon qilinsa boshlang'ich qiymatlari
	//  0 bilan to'ldiriladi
	fmt.Println(a) // [0 0 0]

	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3} //... qo'yilsa uzunligini go o'zi belgilarydi
	d := [3]int{1: 12} //1-indeksga qiymat sifatida 12 ni
	//  yozadi va uzunligi n+1 bo'ladi

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12 0]
}

func arrayCheck() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{3, 2, 1}

	// arrayni qiymatlarini solishtirishi yo'li
	// array uzunligi va qiymatlari tipi bir xil bo'lishi shart
	fmt.Println(a == b) // true
	fmt.Println(a == c) // false
}

func main() {
	arrayDeclered()
	arrayCheck()
}

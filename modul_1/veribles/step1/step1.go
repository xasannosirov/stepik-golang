// 1.5

package main

import "fmt"

// o'zgaruvchilarni e'lon qilish
func VeribleDecelered() {
	var num1 int
	var num2, num3, num4 int

	fmt.Println(num1, num2, num3, num4) //0 0 0 0
}

// o'zgaruvchilarni e'lon qilish va ularga qiymat berish
func VeribleValue() {
	var hello string
	hello = "Hello Go!"
	var a int = 2019

	fmt.Println(hello)
	fmt.Println(a)

	var symbol int32 = 'c'
	fmt.Println(string(symbol))
}

func main() {
	VeribleDecelered()
	VeribleValue()
}

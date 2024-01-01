package main

import (
	"fmt"
	"strconv"
)

func IntToString() {
	a := strconv.Itoa(2020) // int -> string
	fmt.Printf("%T \n", a)  // type - string
	fmt.Println(a)          // 2020
}

func IntToStringConv() {
	user := "talabasi"
	steps := 4
	fmt.Println("Tabriklamiz siz", strconv.Itoa(steps), "-kurs", user, "bo'ldingiz")
	// Tabriklaymiz siz 4 -kurs talabasi bo'ldingiz
}

func Formats() {
	var a = true
	res := strconv.FormatBool(a)
	fmt.Println(res)      // true
	fmt.Printf("%T", res) // string

	var b float64 = 1.0123456789
	fmt.Println(strconv.FormatFloat(b, 'f', 2, 64))  // 1.01
	fmt.Println(strconv.FormatFloat(b, 'f', -1, 64)) // 1.0123456789

	var c int64 = 0xB                     // 16 lik sanoq sistemasidagi B 11 ga teng
	fmt.Println(strconv.FormatInt(c, 10)) // 11
}

func main() {
	IntToString()
	IntToStringConv()
	Formats()
}

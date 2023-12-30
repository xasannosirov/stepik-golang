package main

import (
	"fmt"
)

// fmt.Print() and fmt.Println() ning farqi
func PrintType() {
	fmt.Print("hello, world")
	fmt.Print("hello, world")
	// hello, worldhello, world
	fmt.Print("hello, world")
	fmt.Print("hello, world")
	// hello, world
	// hello, world
	fmt.Print("Ivan", 27)   // Ivan27
	fmt.Println("Ivan", 27) // Ivan 27
	fmt.Print(33, 27)       // 33 27
}

// so'zlarni boshqa so'zlarga qo'shib bir yaxlit ko'nishda chiqarish
func MergeWords() {
	name := "Ivan"
	age := 27
	fmt.Println("My name is", name, "and I am", age, "years old.")
	// My name is Ivan and I am 27 years old.
}

func main(){
	PrintType()
	MergeWords()
}

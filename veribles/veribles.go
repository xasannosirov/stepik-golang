// 1.5

package main

import "fmt"

func myFunc1() {
	// o'zgaruvchilarni e'lon qilish
	var num1 int
	var num2, num3, num4 int
	fmt.Println(num1, num2, num3, num4) //0 0 0 0
}

func myFunc2() {
	// o'zgaruvchilarni e'lon qilish va ularga qiymat berish
	number1 := 5
	var number2 int = 5
	var number3 = 5
	fmt.Println(number1, number2, number3) //5 5 5
}

func myFunc3() {
	//declered veribles with var keyword
	var (
		name string = "Dima"
		age  int    = 23
	)
	fmt.Println(name, age) //Dima 23
}

func myFunc4() {
	a := 100
	b := 10
	c1 := a + b                 // с = 110
	c2 := a * b                 // с = 1000
	c3 := a - b                 // с = 90
	c4 := a / b                 // с = 10
	fmt.Println(c1, c2, c3, c4) //110 1000 90 10
}

func myFunc5() {
	fmt.Print("Ivan", 27)   // Ivan27
	fmt.Println("Ivan", 27) // Ivan 27
	fmt.Print(33, 27)       // 33 27
}

func myFunc6() {
	name := "Ivan"
	age := 27
	fmt.Println("My name is", name, "and I am", age, "years old.")
}

//1.6

/*
this is
multiline
comment
*/

func main() {
	myFunc1()
	myFunc2()
	myFunc3()
	myFunc4()
	myFunc5()
	myFunc6()
}

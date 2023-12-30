package main

import "fmt"

// fmt.Scan() yordamida terminaldan ma'lumotlarni olish
func main() {
	var name string
	var age int
	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	fmt.Println(name, age)
}

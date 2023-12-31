package main

import "fmt"

func main() {
	var age, name = add(4, 5, "Tom", "Simpson")
	fmt.Println(age)  // 9
	fmt.Println(name) // Tom Simpson
}

// ikkita int va string ma'lumotni qo'shib qaytaradi
func add(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

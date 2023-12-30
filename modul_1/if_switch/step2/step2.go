package main

import "fmt"

// if shartida bir necha holatlarni tekshirish
func main() {
	a, b := 6,7
	if a < b {
		fmt.Println("a b dan kichik")
	} else if a > b {
		fmt.Println("a b dan katta")
	} else {
		fmt.Println("a b ga teng")
	}
}

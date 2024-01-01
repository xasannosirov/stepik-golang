package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ExampleFunctionWithoutName() {
	src := "aBcDeFg"
	test := "AbCdEfG"
	
	// Map() funksiyasini ichiga anonim funcksiya yozilgan 
	src = strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)
	}, src)

	fmt.Printf("Teskari string: %s. Natija: %v.\n", src, src == test)
}

func Example() {
	// funksiya chaqirilib qiymat berilganida ishga tushadi
	fn := func(a, b int) int { return a + b }

	// shu joyida ishga tushadi
	func(a, b int) {
		fmt.Println(a + b)
	}(12, 34)

	fmt.Println(fn(17, 15))
}

func main() {
	ExampleFunctionWithoutName()
	Example()
}

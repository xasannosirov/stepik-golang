package main

import (
	"fmt"
	"strings"
	"unicode"
)

// katta belgini kichik, kichigini katta qiladi
func invert(r rune) rune {
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	return unicode.ToLower(r)
}

func ExampleFirstClassFunctionArgument() {
	src := "aBcDeFg"
	test := "AbCdEfG"
	// Map() funksiyasiga funksiya berilishi
	src = strings.Map(invert, src)

	fmt.Printf("Teskari string: %s. Natija: %v.\n", src, src == test)
}

func main(){
	ExampleFirstClassFunctionArgument()
}

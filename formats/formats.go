package main

import "fmt"

func func_sprintf() {
	var a float64 = 100.123456789
	result := fmt.Sprintf("%.2f", a)
	fmt.Printf("%q\n", result) // "100.12"
}

func types() {
	var a rune = 'Ы'
	fmt.Printf("%q\n", a) //'Ы'

	var a1 byte = 's'
	var a2 int = 1234
	fmt.Printf("%q %b\n", a1, a2) //'s' 10011010010

	fmt.Println("This string\nspans multiple\nlines.")
	// This string
	// spans multiple
	// lines.
}

func main() {
	func_sprintf()
	types()
}

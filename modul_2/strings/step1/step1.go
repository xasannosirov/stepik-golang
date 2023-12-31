package main

import "fmt"

func main() {
	ExampleString()
}

// stringlar ustida amallar bajarish
func ExampleString() {
	var s string = "This is string"
	fmt.Printf("String uzunlig: %d bayt\n", len(s))
	fmt.Printf("This is second word: \"%v\"\n", s[7:])
	s = s + "New string"
	fmt.Printf("%v\n", s)
	for _, b := range s {
		fmt.Printf("%v ", b)
	}
	fmt.Print("\n")
}

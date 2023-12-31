package main

import "fmt"

func main() {
	ExampleString()
}

// stringlar ustida amallar bajarish
func ExampleString() {
	var s string = "Это строка"
	fmt.Printf("Длина строки: %d байт\n", len(s))
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\"\n", s[7:])
	s = s + " Новая строка"
	fmt.Printf("%v\n", s)
	for _, b := range s {
		fmt.Printf("%v ", b)
	}
	fmt.Print("\n")
}

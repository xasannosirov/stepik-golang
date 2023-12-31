package main

import "fmt"

func main() {
	ExampleRune()
}

// rune ni indeki orqali o'zgratirish
func ExampleRune() {
	rs := []rune("This slice rune")

	for i := range rs {
		if rs[i] == 's' {
			rs[i] = '*'
		}
	}
	fmt.Printf("Result: %s\n", string(rs))
}

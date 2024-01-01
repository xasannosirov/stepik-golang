package main

import (
	"fmt"
)

func main() {
	a := "str"

	b := []byte(a)

	c := string(b)

	fmt.Println(a) // str

	fmt.Println(b) // [115 116 114] - slice byte

	fmt.Println(c) // str
}

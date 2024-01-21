package main

import (
	"fmt"
	"unicode"
)

func IsStrongPassword(r []rune) string {
	if len(r) < 5 {
		return "Wrong password"
	}
	for i := range r {
		if !(unicode.Is(unicode.Latin, r[i]) || unicode.IsDigit(r[i])) {
			return "Wrong password"
		}
	}
	return "Ok"
}

func main() {
	var str string
	fmt.Scan(&str)
	r := []rune(str)
	fmt.Println(IsStrongPassword(r))
}

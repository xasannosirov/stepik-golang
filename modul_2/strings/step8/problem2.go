package main

import (
	"fmt"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func polindrom(text string) string {
	if text == reverse(text) {
		return "Палиндром"
	} else {
		return "Нет"
	}
}

func main() {
	var text string
	fmt.Scan(&text)

	fmt.Println(polindrom(text))
}

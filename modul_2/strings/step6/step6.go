package main

import (
	"fmt"
	"unicode/utf8"
)

// RuneCountInString() funcksiyasi yordamida inglizcha va ruscha 
// berlgilarni uzunligini sanashda ishlamiz
func main() {
	var en = "english"
	var ru = "русский"
	fmt.Println(len(en), len(ru))
	fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru))
}

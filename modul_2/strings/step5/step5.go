package main

import (
	"fmt"
	"unicode"
)

// stringlar ustida amallar bajarayotganda unicode packageni
// import qilib u yerdagi tayyir funksiyalardan foydalanishimiz mumkin
func main() {
	fmt.Println(unicode.IsDigit('1')) // true
	fmt.Println(unicode.IsLetter('a')) // true 
	fmt.Println(unicode.IsLower('A')) // false
	fmt.Println(unicode.IsUpper('A')) // true
	fmt.Println(unicode.IsSpace('\t')) // true 
 	fmt.Println(unicode.Is(unicode.Latin, 'ы')) // false
	fmt.Println(string(unicode.ToLower('F'))) // f
	fmt.Println(string(unicode.ToUpper('f'))) // F
}
package main

import "fmt"

// b pointeri orqali a o'zgaruvchisini qiymatini o'zgatiramiz
func main() {
	a := 200
	b := &a
	*b++
	fmt.Println(a)
}

package main

import "fmt"

func main() {
	hello() //funksiya chaqirildi
}

// funksiya hechnarsa qabul qilmaydi va qaytarmaydi, lekin konsolga chiqaradi
func hello() {
	fmt.Println("Hello World")
}

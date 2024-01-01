package main

import "fmt"

// switch case yordamida qanaqa tipdagi ma'lumot kelganiga qarab 
// har xil amal bajarishi
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("2 ga ko'paytirildi:", v*2)
	case string:
		fmt.Println(v + " golang")
	default:
		fmt.Printf("Men tushunmaydigan tip --> %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

package main

import "fmt"

func ChangeValue() {
	a := 100
	b := 10
	// bir o'zgaruvchini har xil qiymatga o'zarishi
	c := a + b // с = 110
	fmt.Println(c)
	c = a * b // с = 1000
	fmt.Println(c)
	c = a - b // с = 90
	fmt.Println(c)
	c = a / b // с = 10
	fmt.Println(c)
}

// o'zgaruvchiga ma'lum bir amallarning natijalarini tenglash
func ResultValue() {
	var a int = 10 / 6
	var m float32 = 10.0 / 6
	var c int = 10 % 3
	var d int = 1
	d++
	fmt.Printf("%d\n%f\n%d\n%d\n",a,m,c,d)
}

func main(){
	ChangeValue()
	ResultValue()
}

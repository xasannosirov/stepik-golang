package main

import (
	"fmt"
	"math"
)

// Shape nomli interface
type Shape interface {
	Area() float64
}

// Rectangle nomli struct
type Rectangle struct {
	Width, Height float64
}

// Circle nomli struct
type Circle struct {
	Radius float64
}

// Rectangle structiga tegishli Area() methodi
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle structiga tegishli Area() methodi
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

/* 
getArea() nomli funksiya

Circle yoki Rectangle obyekti kelishiga qarab har xil ishlaydi
yani o'zini Area() methodini ishga tushirib beradi
*/
func getArea(shape Shape) {

	fmt.Println(shape.Area())
}

func main() {

	r := Rectangle{Width: 7, Height: 8}
	c := Circle{Radius: 5}

	getArea(r)
	getArea(c)
}

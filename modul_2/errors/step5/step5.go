package main

import "fmt"

func main() {
	fmt.Println(divide(15, 5))
	fmt.Println(divide(4, 0))
	fmt.Println("Program has been finished")
}
func divide(x, y float64) float64 {
	if y == 0 {
		panic("division by zero!") //dasturni to'xtatishda ishlatiladi
	}
	return x / y
}

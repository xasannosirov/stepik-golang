package main

import "fmt"

func IntToFloat() {
	var x int64 = 57
	var y float64 = float64(x)
	fmt.Print(y) // 57
}

func FloatToInt() {
	var f float64 = 56.231
	var i int = int(f)
	fmt.Println(f) // 56.231
	fmt.Println(i) // 56
}

func FloatAndInt() {
	a := 5 / 2
	fmt.Println(a) // 2

	b := 5.0 / 2
	fmt.Println(b) //2.5
}

func main(){
	IntToFloat()
	FloatToInt()
	FloatAndInt()
}

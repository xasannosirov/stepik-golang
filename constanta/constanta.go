// 1.7

package main

import "fmt"

func exampleConst() {
	// declered const and output type in consol
	const pi float64 = 3.1415
	fmt.Printf("%T\n", pi)

	// cdeclered const and output value in conson
	const (
		a int     = 45
		b float32 = 3.3
	)
	fmt.Println(a, b)

	// get the value after itself
	const (
		A int = 45
		B
		C float32 = 3.3
		D
	)
	fmt.Println(A, B, C, D)

	// increase in value
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_ // 7 skip
		Add
	)
	fmt.Println(Sunday)   //  0
	fmt.Println(Saturday) //  6
	fmt.Println(Add)      //  8
}

func main() {
	exampleConst()
}

package main

import "fmt"

func main() {
	type Month int

	// oylarga raqamlarni ketma-ket berish
	const (
		January Month = 1 + iota
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)
	fmt.Println(January)
	fmt.Println(March)
	fmt.Println(August)
	fmt.Println(December)

}

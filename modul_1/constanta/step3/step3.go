package main

import "fmt"

// iota kalit so'zi bilan keyingi qiymatlarni doimiy oshishini taminlash
func example1() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday)   // 0
	fmt.Println(Saturday) // 6
}

// iota e'lon qilinsa keyingi iota e'lon qilinsa ham oshishi
func example2() {
	const (
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)
	fmt.Println(c0, c1, c2) // вывод: 0 1 2
}

// _ belgi bilan qiymatni o'tkazib yuborish
func example3() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_ // 7 qiymati o'tkazib yuborildi
		Add
	)

	fmt.Println(Sunday)   // 0
	fmt.Println(Saturday) // 6
	fmt.Println(Add)      // 8
}

func main(){
	example1()
	example2()
	example3()
}

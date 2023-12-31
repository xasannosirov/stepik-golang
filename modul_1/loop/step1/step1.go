package main

import "fmt"

// 0 dan 10 gacga sonlarni yig'indisini topish
func for1() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 45
}

// 0 dan 10 gacga sonlarni kvadaratini topish 1-usuli
func for2() {
	var i = 1
	for ; i < 10; i++ {
		fmt.Println(i * i) // 1, 4, 9, 16, 25, 36, 49, 64, 81
	}
}

// 0 dan 10 gacga sonlarni kvadaratini topish 2-usuli
func for3() {
	var i = 1
	for i < 10 {
		fmt.Println(i * i) // 1, 4, 9, 16, 25, 36, 49, 64, 81
		i++
	}
}

// 0 dan 10 gacga sonlarni kvadaratini topish 3-usuli
func for4() {
	var i = 1
	for i < 10 {
		fmt.Println(i * i) // 1, 4, 9, 16, 25, 36, 49, 64, 81
		i++
	}
}

func for5() {
	var n int
	// 0 kiritmagunicha davom etadi
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)
	}
}

func main() {
	for1()
	for2()
	for3()
	for4()
	for5()
}

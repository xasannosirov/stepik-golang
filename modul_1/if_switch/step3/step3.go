package main

import "fmt"

// ko'p holatlarni if da tekshirish
func if1() {
	i := 4
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else if i == 4 {
		fmt.Println("Four")
	} else if i == 5 {
		fmt.Println("Five")
	}
}

// ko'p holatlarni switch case yordamida tekshirish
func switch1() {
	i := 3
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}

// fallthrough yodamida keyingi case ga tekshirmasdan o'tishi
func switch2() {
	v := 42
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}
}

// switch case bilan bir necha xil shartlarni tekshirish
func switch3() {
	var c uint32
	fmt.Scan(&c)
	switch {
	case 1 <= c && c <= 9:
		fmt.Println("1 dan 9 gacha")
	case 100 <= c && c <= 250:
		fmt.Println("100 dan 250 gacha")
	case 1000 <= c && c <= 6000:
		fmt.Println("1000 dan 6000 gacha")
	}
}

func main() {
	if1()
	switch1()
	switch2()
	switch3()
}

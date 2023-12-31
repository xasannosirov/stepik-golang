package main

import (
	"errors"
	"fmt"
)

// bir sonni ikkinchisiga bo'lib natijasini qaytaradigan funksiya
// agar bo'lish mukin bo'lmagan holatlar bo'lsa error qaytaradi
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ZeroDivisionError")
	} else {
		return a/b, nil
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(result)
	}
}

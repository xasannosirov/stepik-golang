package main

import (
	"errors"
	"fmt"
)

// errors.New() yordamida yangi error yaratishimiz mumkin
// fmt.Errorf() yodamida ham qilsa bo'aldi

func main() {
	err := errors.New("my error")
	fmt.Println("", err)
}

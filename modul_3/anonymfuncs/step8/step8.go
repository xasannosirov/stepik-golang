package main

import (
	"errors"
	"fmt"
)

// birinchi kirgan defer funksiya oxirda ishlaydi
func ExampleDefer1() {
	defer func() { fmt.Println(1) }()

	defer func() { fmt.Println(2) }()

	defer func() { fmt.Println(3) }()
}

// defer yordamida anomim funksiyani ishga tushiradi
func someFuncWithPanic() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	// dasturni "fatal error" bilan to'xtatishda ishlatiladi
	panic(errors.New("fatal error"))
}

// agagr someFuncWithPanic funksiyasidan error qaytsa shu errorni chiqaradi
func ExamplePanicRecover() {
	if err := someFuncWithPanic(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	ExampleDefer1()
	ExamplePanicRecover()
}

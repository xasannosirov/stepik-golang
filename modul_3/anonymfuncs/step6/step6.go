package main

import "fmt"

// print qiladigan funksiyani return qiladi
func externalFunction() func() {
	text := "TEXT"

	return func() {
		fmt.Println(text)
	}
}

// return qaytaradigan funksiya chaqirilganda u ishga tushib print qiladi
func ExampleEnvironment() {
	fn := externalFunction()
	fn()
}

// birdan 5 gacha sonlarni kvadratrini topadigan funksiya
func ExampleClosure() {
	fn := func() func(int) int {
		count := 0
		return func(i int) int {
			count++
			return count * i
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println(fn(i))
	}
}

func main() {
	ExampleEnvironment()
	ExampleClosure()
}

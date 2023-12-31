package main

import "fmt"

func main() {
	// defer kalit so'zi yordamida vazifalarni oxirida bajarilishini taminlash mukin
	defer finish()
	defer fmt.Println("Program has been started")
	fmt.Println("Program is working")
}

func finish() {
	fmt.Println("Program has been finished")
}

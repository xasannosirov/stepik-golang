package main

import "fmt"

// o'zgaruvchini qiymatini 1 ga o'zgartirib qo'yadi
func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}

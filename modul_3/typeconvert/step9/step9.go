package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "23.23456"
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)         // 23.23456
	fmt.Printf("%T \n", result) // float64

	s = "1.0000000012345678"
	result32, _ := strconv.ParseFloat(s, 32)
	result64, _ := strconv.ParseFloat(s, 64)
	fmt.Println("bitSize32:", result32) // 1
	fmt.Println("bitSize64:", result64) // 1.0000000012345678
}

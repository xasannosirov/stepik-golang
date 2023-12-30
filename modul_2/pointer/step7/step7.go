package main

import "fmt"

func main() {
	a := 200
	b := &a
	*b++
	c := &b
	**c++ // pointerga pointer
	fmt.Println(a)
}

package main

import "fmt"

func myPrint(a ...interface{}) {
	for _, elem := range a {
		fmt.Printf("%d ", elem)
	}
}

func ExampleMyPrint() {
	myPrint(1, 2, 3, 4, 5)
}

func ExampleExpandSlice1() {
	s := []interface{}{1, 2, 3, 4, 5}

	fmt.Println(s)
	fmt.Println(s...)
}

func ExampleExpandSlice2() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}

	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func main() {
	ExampleMyPrint()
	ExampleExpandSlice1()
	ExampleExpandSlice2()
}

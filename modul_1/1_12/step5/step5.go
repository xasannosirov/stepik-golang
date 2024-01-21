package main

import "fmt"

func main() {
	var workArray = [10]uint8{}

	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	var a, b int
	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}
	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
}

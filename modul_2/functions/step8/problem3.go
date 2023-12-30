package main

import "fmt"

func vote(x int, y int, z int) int {
	if x+y == 0 {
		return 0
	} else if x+y == 1 && z == 0 {
		return 0
	} else {
		return 1
	}
}

func main() {
	fmt.Println(vote(0, 0, 1))
}

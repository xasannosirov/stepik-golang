package main

import (
	"fmt"
	"time"
)

func work(n int) int {
	time.Sleep(time.Second)
	if n > 3 {
		return n + 1
	}
	return n - 1
}

func main() {
	m := make(map[int]int)
	var a int
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		if m[a] == 0 {
			m[a] = work(a)
		}
		fmt.Print(m[a], " ")
	}
	fmt.Println("time limit ok")
}

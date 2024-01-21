package main

import (
	"fmt"
)

func main() {
	var str, res string
	fmt.Scan(&str)
	r := []rune(str)
	for i := range r {
		if i%2 == 1 {
			res += string(r[i])
		}
	}
	fmt.Println(string(res))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, res string
	fmt.Scan(&str)
	r := []rune(str)
	for i := range r {
		if strings.Count(string(r), string(r[i])) == 1 {
			res += string(r[i])
		}
	}
	fmt.Println(string(res))
}

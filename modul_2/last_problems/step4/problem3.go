package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	max := int(str[0])
	for i := range str {
		if max < int(str[i]) {
			max = int(str[i])
		}
	}
	fmt.Println(max - 48)
}

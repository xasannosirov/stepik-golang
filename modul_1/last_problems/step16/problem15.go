package main

import "fmt"

func main() {
	var n, m, temp string
	fmt.Scan(&n, &m)
	for _, elem := range n {
		if string(elem) != m {
			temp += string(elem)
		}
	}
	fmt.Println(temp)
}

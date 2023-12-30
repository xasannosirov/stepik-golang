package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	for _, e := range a {
		if strings.Index(b, string(e)) != -1 {
			fmt.Print(string(e), " ")
		}
	}
}
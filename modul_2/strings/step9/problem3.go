package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(strings.Index(a, b))
}

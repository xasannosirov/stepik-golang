package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	x := strings.Replace(str, "", "*", -1)
	fmt.Println(x[1 : len(x)-1])
}

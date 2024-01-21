package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)
	r := strconv.Itoa(num)
	res := ""
	for _, i := range r {
		res += strconv.Itoa((int(i) - 48) * (int(i) - 48))
	}
	e, _ := strconv.Atoi(res)
	fmt.Println(e)
}

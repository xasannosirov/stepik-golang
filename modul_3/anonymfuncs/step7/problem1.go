package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(a uint) uint {
		var s, z string = strconv.FormatUint(uint64(a), 10), ""
		for _, v := range s {
			i, _ := strconv.Atoi(string(v))
			if i != 0 && i%2 == 0 {
				z += strconv.Itoa(i)
			}
		}
		g, _ := strconv.Atoi(z)
		if g == 0 {
			return uint(100)
		} else {
			return uint(g)
		}
	}
	fmt.Println(fn(727178))
}

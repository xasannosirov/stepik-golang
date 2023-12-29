package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Scan(&d)
	fmt.Println("It is", d/30, "hours", 2*(d%30), "minutes.")
}

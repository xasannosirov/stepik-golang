package main

import (
	"fmt"
	"time"
)

func main() {
	var input string
	fmt.Scan(&input)
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return
	}
	fmt.Println(t.Format(time.UnixDate))
}

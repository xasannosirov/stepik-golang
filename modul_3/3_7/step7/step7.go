package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	text := strings.Split(in.Text(), ",")
	clockFuture, _ := time.Parse("02.01.2006 15:04:05", text[0])
	clockPast, _ := time.Parse("02.01.2006 15:04:05", text[1])
	if clockFuture.After(clockPast) {
		fmt.Println(clockFuture.Sub(clockPast))
	} else {
		fmt.Println(clockPast.Sub(clockFuture))
	}
}

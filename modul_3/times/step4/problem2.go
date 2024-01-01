package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, err := time.Parse("2006-01-02 15:04:05", scanner.Text())
	if err != nil {
		panic(err)
	}
	if t.Hour() >= 13 {
		t = t.Add(24 * time.Hour)
	}
	fmt.Println(t.Format("2006-01-02 15:04:05"))

}

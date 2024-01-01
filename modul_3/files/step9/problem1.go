package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			os.Exit(0)
		}
		sum += i
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}

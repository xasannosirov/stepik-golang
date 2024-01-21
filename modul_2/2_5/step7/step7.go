package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	r := []rune(text)

	if unicode.IsUpper(r[0]) && string(r[len(r)-1]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

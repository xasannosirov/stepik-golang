package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file) //faylga yozadigan obyekt yasaldi
	n, err := w.WriteString("Thera is data in file") //faylga yozilishi
	if err != nil {
		return
	}
	fmt.Printf("%d byte malumot yozib ikindi\n", n)
	w.Flush()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1 149,6088607594936;1 179,0666666666666
shu ma;lumotni o'qib olib ikkita float son ko'rinishiga keltirib
bo'linmasini topish kerak, terminalda +Inf deb chiqaradi lekin stepik qabul qiladi
*/

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	arr := strings.Split(strings.TrimSuffix(input, "\r\n"), ";")
	var t1, t2 string
	for _, elem := range arr[0] {
		if elem == ' ' {
			t1 += ""
		} else if elem == ',' {
			t1 += "."
		} else {
			t1 += string(elem)
		}
	}
	for _, elem := range arr[1] {
		if elem == ' ' {
			t2 += ""
		} else if elem == ',' {
			t2 += "."
		} else {
			t2 += string(elem)
		}
	}
	t3, _ := strconv.ParseFloat(t1, 64)
	t4, _ := strconv.ParseFloat(t2, 64)

	fmt.Printf("%.4f", t3/t4)
}

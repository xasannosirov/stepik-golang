package main

import (
	"fmt"
)

func main(){
	var n, temp, summ int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		if temp % 8 == 0 && 0 < temp/10 && temp/10 <= 9 {
			summ += temp
		}
	}
	fmt.Println(summ)
}
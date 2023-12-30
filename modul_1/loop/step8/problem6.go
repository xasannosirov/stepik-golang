package main

import (
	"fmt"
)

func main(){
	var n, temp int
	for n < 100 {
		fmt.Scan(&temp)
		if temp < 10 {
			continue
		} else if temp > 100 {
			break
		} else {
			fmt.Println(temp)
		}
		n = temp
	}
}
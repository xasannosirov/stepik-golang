package main

import (
	"fmt"
)

func main(){
	var n, max_num, count_num = -1, -1, 0
	var temp int
	for n != 0 {
		fmt.Scan(&temp)
		n = temp
		if temp == max_num {
			count_num += 1
		} else if temp > max_num {
			count_num = 1
			max_num = temp
		}
	}
	fmt.Println(count_num)
}
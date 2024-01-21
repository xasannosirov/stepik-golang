package main

import "fmt"

func minimumFromFour() int {
	var array [4]int
	for i := 0; i < 4; i++ {
		fmt.Scan(&array[i])
	}
	var min_num = array[0]
	for i := 1; i < 4; i++ {
		if array[i] < min_num {
			min_num = array[i]
		}
	}
	return min_num
}

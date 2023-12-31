package main

import "fmt"

func main() {
	var a, b int
	array := []int{}
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		array = append(array, b)
	}
	var min_num, min_count int = array[0], 0
	for i := 0; i < a; i++ {
		if array[i] == min_num {
			min_count += 1
		} else if array[i] < min_num {
			min_num = array[i]
			min_count = 1
		}
	}
	fmt.Println(min_count)
}

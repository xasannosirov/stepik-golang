package main

import "fmt"

func main() {
	var number float64
	fmt.Scan(&number)

	if number >= 10000 {
		fmt.Println(fmt.Sprintf("%e", number))
	} else if number <= 0 {
		fmt.Println("число", fmt.Sprintf("%.2f", number), "не подходит")
	} else {
		fmt.Println(fmt.Sprintf("%.4f", number*number))
	}
}

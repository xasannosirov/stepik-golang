package main

import "fmt"

func main() {
	var n uint
	fmt.Scan(&n)
	if n%10 == 1 && n != 11 {
		fmt.Println(n, "korova")
	} else if n%10 > 1 && n%10 < 5 && n/10 != 1 {
		fmt.Println(n, "korovy")
	} else {
		fmt.Println(n, "korov")
	}
}

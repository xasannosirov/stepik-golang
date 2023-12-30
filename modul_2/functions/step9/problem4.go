package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
    var f1, f2 int = 1, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f1
}

func main(){
	fmt.Println(fibonacci(4))
}
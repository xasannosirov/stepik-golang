package main

import "fmt"

/*
bu funksiya channel va int tipida son qabul qiladi
va int songa 1 ni qo'shin natijasini kanalga yozib qo'yadi
*/
func task(c chan int, n int) {
	c <- n + 1
}

func main() {
	chan1 := make(chan int, 2)
	task(chan1, 2)
	fmt.Println(<-chan1)
}

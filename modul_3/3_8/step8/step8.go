package main

func task(c chan int, n int) {
	c <- n + 1
}

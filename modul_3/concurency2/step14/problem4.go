package main

import "fmt"

func calculator(argements <-chan int, done <-chan struct{}) <-chan int {
	c := make(chan int)
	sum := 0
	go func() {
		defer close(c)
		for {
			select {
			case arg := <-argements:
				sum += arg
			case <-done:
				c <- sum
				return
			}
		}
	}()
	return c
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})
	result := calculator(arguments, done)
	for i := 0; i < 10; i++ {
		arguments <- 1
	}
	close(done)
	fmt.Println(<-result)
}

package main

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

package main

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		select {
		case arg := <-firstChan:
			c <- arg * arg
		case arg := <-secondChan:
			c <- arg * 3
		case <-stopChan:
			return
		}
	}()
	return c
}

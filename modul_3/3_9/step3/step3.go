package main

import "time"

func work() {
	time.Sleep(time.Second)
}

func main() {
	done := make(chan struct{})
	go func(d chan struct{}) {
		work()
		close(d)
	}(done)
	<-done
}

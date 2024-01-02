package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	wg := new(sync.WaitGroup)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, tick.C, wg)
	}

	wg.Wait()

	/*
	 * worker 1 ishni qildi
	 * worker 5 ishni qildi
	 * worker 3 ishni qildi
	 * worker 4 ishni qildi
	 * worker 2 ishni qildi
	 */
}

func worker(id int, limit <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	<-limit
	fmt.Printf("worker %d ishni qildi\n", id)
}

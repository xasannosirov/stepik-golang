package main

import (
	// "fmt"
	// "time"
	"sync"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var c chan int = make(chan int)
	go func() {
		defer close(c)
		var arr1, arr2 []int
		wg := new(sync.WaitGroup)
		mu := new(sync.Mutex)

		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				mu.Lock()
				defer wg.Done()
				arr1 = append(arr1, <-in1)
				mu.Unlock()
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				mu.Lock()
				defer wg.Done()
				arr2 = append(arr2, <-in2)
				mu.Unlock()
			}(wg)
		}

		wg.Wait()
		for i := 0; i < n; i++ {
			out <- arr1[i] + arr2[i]
		}
	}()
}

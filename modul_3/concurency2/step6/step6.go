package main

import (
	"fmt"
	"sync"
)

// sync.Mutex() yordamida bir nechta gorutina bir vaqtda ishlab
// bir xil qiymat olib ketishini oldini olish
func main() {
	var x int
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock() //boshqa gorutinalarni block qilib turadi
			x++
			mu.Unlock() //blockdan ochadi
		}(wg, mu)
	}

	wg.Wait()
	fmt.Println(x)
}

package main

import (
	"fmt"
	"sync"
)
/*
in1 va in2 channeldan kelgan ma'lumotni o'qib uni fn funksiyaga berib, natijasini
out channelga yozadi, shu ketma-ketlikni n marta bajaradi
*/ 

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var (
		slice1   = make([]int, n)
		slice2   = make([]int, n)
		wg sync.WaitGroup
		mu sync.Mutex
	)
	wg.Add(n * 2)
	// in1 channeldagi ma'lumotlarni o'qib oladi
	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			slice1[i] = <-in1
			mu.Unlock()
			wg.Done()
		}
	}()
	//in2 channeldagi ma'lumotlarni o'qib oladi
	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			slice2[i] = <-in2
			mu.Unlock()
			wg.Done()
		}
	}()
	//in1 channeldan o'qib olgan ma'lumotini fn funksiyaga berib natijasini oladi
	go func() {
		for i, v := range slice1 {
			slice1[i] = fn(v)
		}
	}()
	//in2 channeldan o'qib olgan ma'lumotini fn funksiyaga berib natijasini oladi
	go func() {
		for i, v := range slice2 {
			slice2[i] = fn(v)
		}
	}()
	// ikkala channeldagi ma'lumotlani fn ga berganidan keyin yig'indisni fn ga yozadi
	go func() {
		defer close(out)
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- slice1[i] + slice2[i]
		}
	}()
}

func main() {
	in1 := make(chan int, 6)
	in2 := make(chan int, 6)
	out := make(chan int, 6)

	for i := 0; i <= 5; i++ {
		in1 <- i
		in2 <- i
	}

	fn := func(n int) int {
		return n
	}

	merge2Channels(fn, in1, in2, out, 6)

	for i := 0; i <= 5; i++ {
		fmt.Println(<-out)
	}
}

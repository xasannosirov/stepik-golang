package main

import (
	"fmt"
	"sync"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var (
		slice1 = make([]int, n)
		slice2 = make([]int, n)
		wg     sync.WaitGroup
		mu     sync.Mutex
	)
	wg.Add(n * 2)
	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			slice1[i] = <-in1
			mu.Unlock()
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			slice2[i] = <-in2
			mu.Unlock()
			wg.Done()
		}
	}()
	go func() {
		wg.Wait() // Barcha gorutinalarning tugaganini kutish
		close(out)
	}()
	go func() {
		for i, v := range slice1 {
			slice1[i] = fn(v)
		}
	}()
	go func() {
		for i, v := range slice2 {
			slice2[i] = fn(v)
		}
	}()
	go func() {
		wg.Wait() // Barcha gorutinalarning tugaganini kutish
		for i := 0; i < n; i++ {
			out <- slice1[i] + slice2[i]
		}
	}()
}

func main() {
	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)

	fn := func(x int) int {
		// time.Sleep(time.Second)
		return x
	}

	n := 5 // Masalan, n qiymatini o'zingizga mos ravishda o'zgartiring

	// Funksiyani bajarish uchun gorutina
	go merge2Channels(fn, in1, in2, out, n)

	// Input qiymatlarni yuborish
	go func() {
		for i := 1; i <= n; i++ {
			in1 <- i
			in2 <- i * 2
		}
		close(in1)
		close(in2)
	}()

	// Natijani olish
	for result := range out {
		fmt.Println("Result:", result)
	}
}

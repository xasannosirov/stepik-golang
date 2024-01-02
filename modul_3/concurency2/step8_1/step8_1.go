package main

import (
	"fmt"
	"time"
)

func main() {
	<-work()
	/*
	 * tik-tak
	 * tik-tak
	 * tik-tak
	 * tik-tak
	 */
}

func work() <-chan struct{} {
	done := make(chan struct{}) // goroutinlarni sinxronlashtirish uchun kanal

	go func() {
		defer close(done) // funksiya o'z ishini tugatgandan so'ng sinxronlash kanali yopiladi

		stop := time.NewTimer(time.Second)

		tick := time.NewTicker(time.Millisecond * 200)
		defer tick.Stop() //funktsiya tugashi bilan resurslarni ozod qiling

		for {
			select {
			case <-stop.C:
				// stop - Timer, 1 soniyadan keyin o'chirish uchun signal beradi
				return
			case <-tick.C:
				// tick - Ticker, har 200 millisekundda ishni bajarish uchun signal yuborish
				fmt.Println("tik-tak")
			}
		}
	}()

	return done
}

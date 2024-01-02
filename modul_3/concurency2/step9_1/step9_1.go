package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}

func main() {
	c := make(chan int, 20)
	timeout := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {

		select { // select operator

		case gopherID := <-c: // Goferning uyg'onishini kutish

			fmt.Println("gopher ", gopherID, " has finished sleeping")

		case <-timeout: // Vaqt tugashini kutish

			fmt.Println("my patience ran out")

			return

		}

	}
}

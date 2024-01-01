package main

import "fmt"

// bu funksiya channel va string tipida ma'lumot oladi va
// shu ma'lumotni kanalga 5 marta yozadi
func task2(c chan string, n string) {
	for i := 0; i < 5; i++ {
		c <- n + " "
	}
}

func main() {
	chan1 := make(chan string, 6)
	task2(chan1, "salom")
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
}

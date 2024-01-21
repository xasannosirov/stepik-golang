package main

func task2(c chan string, n string) {
	for i := 0; i < 5; i++ {
		c <- n + " "
	}
}

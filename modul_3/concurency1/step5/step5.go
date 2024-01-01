package main

import "fmt"

func main() {
	channel := make(chan int, 20) //channel yaratish
	
	defer close(channel) //channelni oxirida close qilish

	fmt.Println(len(channel), cap(channel)) // 0 20

	// channelga ma'lumotlarni yozish
	channel <- 5
	channel <- 35
	channel <- 65
	channel <- 57
	channel <- 35
	channel <- 50

	fmt.Println(len(channel), cap(channel)) // 6 20

	// channeldan ma'lumotlarni o'qib olish
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println(len(channel), cap(channel)) // 0 20

}

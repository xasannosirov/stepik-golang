package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(time.Second * 2) // с2 sekund kutib turadi

	timer := time.After(time.Second)
	<-timer // 1 sekunddan keyin channeldan ma'lumot olinadi

	ticker := time.Tick(time.Second)
	count := 0

	for {
		<-ticker //har soniyada channaldan ma'lumot olinadi
		fmt.Println("очередной тик")
		count++
		if count == 3 {
			break
		}
	}
}

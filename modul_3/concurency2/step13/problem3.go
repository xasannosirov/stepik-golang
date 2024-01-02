package main

import (
	"fmt"
	"time"
)

/*
firstChan channeldan ma'lumot kelsa, kelgan ma'lumotni kvadratini c channelga yozadi
secondChan channeldan ma'lumot kelsa, kelgan ma'lumotni 3 ga ko'paytirib c channelga yozadi
stopChan channeldan ma'lumot kelsa, gorutina to'xtaydi, keyin esa c channelni return qiladi
*/ 

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		select {
		case arg := <-firstChan:
			c <- arg * arg
		case arg := <-secondChan:
			c <- arg * 3
		case <-stopChan:
			return
		}
	}()
	return c
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	resultChan := calculator(firstChan, secondChan, stopChan)

	go func() {
		defer close(firstChan)
		defer close(secondChan)
		defer close(stopChan)

		firstChan <- 5
		secondChan <- 7

		time.Sleep(1 * time.Second)

		stopChan <- struct{}{}
	}()

	result := <-resultChan
	fmt.Println("Result:", result)
}

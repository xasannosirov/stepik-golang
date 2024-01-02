package main

import (
	"fmt"
)

func main() {
	// myFunc() uchun
	done := make(chan struct{})
	go myFunc(done)
	<-done //channel close bo'lgunicha kutadi

	//myFunc2() uchun
	<-myFunc2() //close bo'lgan channel kelmaguncha kutadi
}

func myFunc(done chan struct{}) {
	fmt.Println("hello")
	close(done)
}

func myFunc2() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		fmt.Println("hello")
		close(done)
	}()
	return done
}

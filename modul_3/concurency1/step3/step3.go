package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//protsessordago 4 ta yadro ishlashi keraklini kodda ko'rsatish
	runtime.GOMAXPROCS(4)

	go myFunc()                 //bu funksiya maindan tashqari alohida ishlaydi
	time.Sleep(1 * time.Second) // 1 sekund to'xtab turadi
}

func myFunc() {
	fmt.Println("hello")
}

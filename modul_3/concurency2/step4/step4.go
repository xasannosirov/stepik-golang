package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup) 

	for i := 0; i < 5; i++ {
		wg.Add(1) //nechta gorutinani kutishi aytildi
		go work(i, wg) //alohida gorutina sifatida ishga tushirildi
	}

	wg.Wait() //wg.Add() qilib qo'shilgan gorutinani hammasi wg.Done() bo'lgunicha kutadi
	fmt.Println("Gorutinalar bajarishni tugatdi")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done() //gorutina tuganai xaqida waitGroupga xabar beradi
	fmt.Printf("%d-gorutina ishni boshladi\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("%d-gorutina ishni tugatdi\n", id)
}

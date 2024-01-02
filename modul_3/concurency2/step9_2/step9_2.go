package main

import (
	"fmt"
	"time"
)

func main() {
	tick1 := time.After(time.Second)
	tick2 := time.After(time.Second * 2)
	select {
	case <-tick1:
		fmt.Println("Birinchi kanaldan olingan qiymat")
	case <-tick2:
		fmt.Println("Ikkinchi kanaldan olingan qiymat")
	// Standart blok case blokidan oldin bajariladi - Go uchun 1 soniya juda uzoq
	default:
		fmt.Println("Default")
	}
}

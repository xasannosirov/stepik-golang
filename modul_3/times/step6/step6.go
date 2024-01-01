package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	past := now.AddDate(0, 0, -1) // 1 kun ayiradi
	future := now.AddDate(0, 0, 1) // 1 kun qo'shadi

	fmt.Println(time.Since(past).Round(time.Second))   // 24h0m0s
	fmt.Println(time.Until(future).Round(time.Second)) // 24h0m0s

	// stringni timega o'tkazadi
	dur, err := time.ParseDuration("1h12m3s")
	if err != nil {
		panic(err)
	}
	fmt.Println(dur.Round(time.Hour).Hours()) // 1
}

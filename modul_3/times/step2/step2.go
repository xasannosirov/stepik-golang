package main

import (
	"fmt"
	"time"
)

func CurrentDate() {
	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	fmt.Println(current.Date())    // 2020 May 15
	fmt.Println(current.Year())    // 2020
	fmt.Println(current.Month())   // May
	fmt.Println(current.Day())     // 15
	fmt.Println(current.Clock())   // 17 45 12
	fmt.Println(current.Hour())    // 17
	fmt.Println(current.Minute())  // 45
	fmt.Println(current.Second())  // 12
	fmt.Println(current.Unix())    // 1589546712
	fmt.Println(current.Weekday()) // Friday
	fmt.Println(current.YearDay()) // 136
}

func ThisFormat() {
	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	fmt.Println(current.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:12
}

func ChekcTime() {
	firstTime := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

	fmt.Println(firstTime.After(secondTime))  // true
	fmt.Println(firstTime.Before(secondTime)) // false
	fmt.Println(firstTime.Equal(secondTime))  // false
}

func ChangeTime() {
	now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	future := now.Add(time.Hour * 12) // 12 soat qo'shadi
	past := now.AddDate(-1, -2, -3)   // 1yil 2 oy 3 kun ayoradi
	fmt.Println(future.Sub(past))     // 10332h0m0s
}

func main() {
	CurrentDate()
	ThisFormat()
	ChekcTime()
	ChangeTime()
}

package main

import (
	"fmt"
	"time"
)

func FormatTime() {
	now := time.Now()

	currentTime := time.Date(
		2020,     // yil
		time.May, // oy
		15,       // jun
		10,       // soat
		13,       // minut
		12,       // sekund
		45,       // nanasekund
		time.UTC, // timezone
	)

	// https://ru.wikipedia.org/wiki/Unix-%D0%B2%D1%80%D0%B5%D0%BC%D1%8F
	unixTime := time.Unix(
		150000, // sekund
		1,      // nanasekund
	)

	fmt.Println(now.Format("02-01-2006 15:04:05"))         // 15-05-2020 09:58:16
	fmt.Println(currentTime.Format("02-01-2006 15:04:05")) // 15-05-2020 10:13:12
	fmt.Println(unixTime.Format("02-01-2006 15:04:05"))    // 02-01-1970 22:40:00
}

func ParseTime() {
	firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
	if err != nil {
		panic(err)
	}

	// https://www.iana.org/time-zones
	loc, err := time.LoadLocation("Asia/Yekaterinburg")
	if err != nil {
		panic(err)
	}

	secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(firstTime.Format("02-01-2006 15:04:05"))  // 15-05-2020 17:45:00
	fmt.Println(secondTime.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:10
}

func main() {
	FormatTime()
	ParseTime()
}

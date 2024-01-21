package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   {"Texas", "California", "Arizona", "Colorado"},
		100:  {"Washington", "New-York", "Virginia", "Montana"},
		1000: {"Alaska", "Michigan", "Florida", "Georgia"},
	}

	cityPopulation := map[string]int{
		"Texas":      15,
		"California": 25,
		"Arizona":    80,
		"Colorado":   50,
		"Vashington": 400,
		"New-York":   480,
		"Virginia":   670,
		"Montana":    940,
		"Alaska":     1200,
		"Michigan":   3800,
		"Florida":    4500,
		"Jorjiya":    9000,
	}

	for _, value := range groupCity[10] {
		delete(cityPopulation, value)
	}

	for _, value := range groupCity[1000] {
		delete(cityPopulation, value)
	}

	fmt.Println(cityPopulation)
}

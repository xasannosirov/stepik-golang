package main

import "fmt"

func arrayChangeValue() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[0]) // 1
	fmt.Println(numbers[4]) // 5
	numbers[0] = 87 //arrayni qiymatini o'zgartirish
	fmt.Println(numbers[0]) // 87
}

func arrayShowValue() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	// range index va qiymat qaytaradi
	for idx, elem := range a {
		fmt.Printf("%d indexdagi element %d\n", idx, elem)
	}
}

func main(){
	arrayChangeValue()
	arrayShowValue()
}
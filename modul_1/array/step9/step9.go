package main

import "fmt"

func main() {
	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Базовый массив: %v\n", baseArray)

	baseSlice := baseArray[5:8]
	fmt.Printf(
		"Срез, основанный на базовом массиве длиной %d и емкостью %d: %v\n",
		len(baseSlice),
		cap(baseSlice),
		baseSlice,
	)

	pointer := fmt.Sprintf("%p", baseSlice)

	baseSlice = append(baseSlice, 10)
	fmt.Printf("Массив: %v\n", baseArray)
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

	// Output:
	// Массив: [0 1 2 3 4 5 6 7 10 9]
	// Срез длиной 4 и емкостью 5: [5 6 7 10]
	// true

	baseSlice = append(baseSlice, 11, 12, 13)
	fmt.Printf("Массив: %v\n", baseArray)
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

	// Output:
	// Массив: [0 1 2 3 4 5 6 7 10 9]
	// Срез длиной 7 и емкостью 10: [5 6 7 10 11 12 13]
	// false
}

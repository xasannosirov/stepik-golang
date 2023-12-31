package main

import "fmt"

func main() {
	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Asosiy array: %v\n", baseArray)

	baseSlice := baseArray[5:8]
	fmt.Printf(
		"Sliceni uzunli %d va uni sig'imi %d: %v\n",
		len(baseSlice),
		cap(baseSlice),
		baseSlice,
	)

	pointer := fmt.Sprintf("%p", baseSlice)

	baseSlice = append(baseSlice, 10)
	fmt.Printf("Array: %v\n", baseArray)
	fmt.Printf("Slice uzunligi %d va uni sig'imi %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

	baseSlice = append(baseSlice, 11, 12, 13)
	fmt.Printf("Array: %v\n", baseArray)
	fmt.Printf("Slice uzunligi %d va uni sig'imi %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))
}

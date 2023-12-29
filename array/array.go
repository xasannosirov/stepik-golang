package main

import "fmt"

func arrayDeclered() {
	var a [3]int
	fmt.Println(a) // [0 0 0]

	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	d := [3]int{1: 12}

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12 0]
}

func arrayCheck() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{3, 2, 1}

	fmt.Println(a == b) // true
	fmt.Println(a == c) // false
}

func arrayChangeValue() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[0]) // 1
	fmt.Println(numbers[4]) // 5
	numbers[0] = 87
	fmt.Println(numbers[0]) // 87
}

func arrayShowValue() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for idx, elem := range a {
		fmt.Printf("%d indexdagi element %d\n", idx, elem)
	}
}

func sliceDeclered() {
	a := make([]int, 10) //make([]T, length, capacity)
	fmt.Println(a)           // [0 0 0 0 0 0 0 0 0 0]
}

func sliceShowElement() {
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	// базовый массив
	users1 := initialUsers[2:6] // 3-elementdan 6-elementgacha
	users2 := initialUsers[:4]  // 1-element 4-elementgacha
	users3 := initialUsers[3:]  // 4-elementdan oxirigacha
	fmt.Println(users1)         // [Kate Sam Tom Paul]
	fmt.Println(users2)         // [Bob Alice Kate Sam]
	fmt.Println(users3)         // [Sam Tom Paul Mike Robert]
}
func sliceAddElement() {
	a := []int{1, 2, 3}
	a = append(a, 4, 5)
	fmt.Println(a) // [1 2 3 4 5]
}

func sliceRemoveElement() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a) // [1 2 4 5 6 7]
}

func sliceCopyElement() {
	a := []int{1, 2, 3}
	b := make([]int, 3)
	n := copy(b, a)
	fmt.Printf("a = %v\n", a)                  // a = [1 2 3]
	fmt.Printf("b = %v\n", b)                  // b = [1 2 3]
	fmt.Printf("Скопировано %d элемента\n", n) // 3 ta element bor
}

func main() {
	arrayDeclered()
	arrayCheck()
	arrayChangeValue()
	arrayShowValue()
	sliceDeclered()
	sliceShowElement()
	sliceAddElement()
	sliceRemoveElement()
	sliceCopyElement()
}

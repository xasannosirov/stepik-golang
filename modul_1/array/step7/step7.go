package main

import "fmt"

func main() {
	//arrayni turli xil uzunlikda kesib olish
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6] // 3-elementdan 6-elementgacha
	users2 := initialUsers[:4]  // 1-element 4-elementgacha
	users3 := initialUsers[3:]  // 4-elementdan oxirigacha
	fmt.Println(users1)         // [Kate Sam Tom Paul]
	fmt.Println(users2)         // [Bob Alice Kate Sam]
	fmt.Println(users3)         // [Sam Tom Paul Mike Robert]
}

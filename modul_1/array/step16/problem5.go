package main
import "fmt"

func main()  {
	var a,b int
	array := [] int{}
	fmt.Scan(&a)
	for i:= 0; i < a; i++ {
		fmt.Scan(&b)
		array = append(array, b)
	}
	var count int
	for i := 0; i < a; i++ {
		if array[i] >= 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
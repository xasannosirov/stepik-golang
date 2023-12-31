package main

import "fmt"

func main(){
	ExampleByteSlice()
}

func ExampleByteSlice() {
	bs := []byte("This is rune date")

	fmt.Printf("Rune: %v\n", bs)

	for i := range bs {
		if bs[i]%2 == 0 {
			bs[i] = bs[i] + 1
			continue
		}
		bs[i] = bs[i] - 1
	}

	fmt.Printf("Result: %s", bs)
}

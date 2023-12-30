package main

import "fmt"

func ExampleScope1() {
	var v int = 1

	{
		var v string = "2"
		fmt.Println(v)
	}

	fmt.Println(v)
}

func main(){
	ExampleScope1()
}
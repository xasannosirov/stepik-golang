package main

import "fmt"

func ExampleScope1() {
	var v int = 1

	// jingalak qavs bilan alohida blok yaratishimiz mumkin
	// bu blokning o'zgaruvchilari va shuni ichida ishlaydi
	{
		var v string = "2"
		fmt.Println(v)
	}

	fmt.Println(v)
}

func main(){
	ExampleScope1()
}
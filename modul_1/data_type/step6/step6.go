package main

import "fmt"

// logig operatorlar
func logictypes() {
	fmt.Println(true && false) //false
	fmt.Println(true || false) //true
	fmt.Println(!true)         // false
	fmt.Println(!false)        // true
}

func main(){
	logictypes()
}
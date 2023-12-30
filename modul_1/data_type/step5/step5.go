package main

import "fmt"

// string
func StringTypes() {
	fmt.Println("Hello Go"[0])         // 72
	fmt.Println(string("Hello Go"[0])) // H
}

func main(){
	StringTypes()
}
// 1.4  all steps

package main

import "fmt"

func datatypes() {
	numberInt := 456
	numberFloat := 34.5
	numberString := "salom"
	// %T o'zgaruvchinig tipini ko'rsatadi
	fmt.Printf("'%d' sonning tipi '%T'\n", numberInt, numberInt)
	fmt.Printf("'%f' sonning tipi '%T'\n", numberFloat, numberFloat)
	fmt.Printf("'%s' sonning tipi '%T'\n", numberString, numberString)
}

func logictypes(){
	fmt.Println(true && false) //false
	fmt.Println(true || false) //true
	fmt.Println(!true) // false
	fmt.Println(!false) // true
}

func main() {
	datatypes()
	logictypes()
}

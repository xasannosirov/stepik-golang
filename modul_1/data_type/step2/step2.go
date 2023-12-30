package main

import "fmt"

// unsignet int sonlar
func UIntTypes() {
	var num1 uint8 = 255                   //uint8 (from 0- to 255)
	var num2 uint16 = 65535                //uint16 (from 0- to 65535)
	var num3 uint32 = 4294967295           //uint32 (from 0- to 4294967295)
	var num4 uint64 = 18446744073709551615 //uint64 (from 0- to 18446744073709551615)

	fmt.Printf("%T\n%T\n%T\n%T", num1, num2, num3, num4)
}

// int numbers
func IntTypes() {
	var num1 int8 = 127                  //uint8 (from -128 to 127)
	var num2 int16 = 32767               //uint16 (from -32768 tо 32767)
	var num3 int32 = 2147483647          //uint32 (from -2147483648 tо 2147483647)
	var num4 int64 = 9223372036854775807 //uint64 (from -9223372036854775808 tо 9223372036854775807)

	fmt.Printf("%T\n%T\n%T\n%T", num1, num2, num3, num4)
}

func main(){
	UIntTypes()
	IntTypes()
}
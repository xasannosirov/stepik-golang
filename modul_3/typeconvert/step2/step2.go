package main

import "fmt"

func typeConvert() {
	// bit tipdagi sonni bashqa tipga o'tkazish
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)

	fmt.Println(bigIndex)         // 15
	fmt.Printf("%T \n", bigIndex) // int32
}

func typeConvertWithoutvalue() {
	// bir tipdan boshqa tipg o'kazishda qiymatning yo'qolishi
	var big int64 = 129
	var little = int8(big)
	fmt.Println(little) //-127
}

func main(){
	typeConvert()
	typeConvertWithoutvalue()
}
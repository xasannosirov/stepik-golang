package main

import (
	"fmt"
	"os"
)

func main() {
	file1, err1 := os.Create("text.txt") //fayl yaratib uni ochadi
	if err1 != nil {
		fmt.Println(err1) //yaratib ochishda muammo bo'lsa uni chiqaradi va to'xtaydi
		return
	}
	os.Rename("text.txt", "newfile.txt") //fayl nomini o'zgartiradi
	file1.WriteString("this is file data1\n") //faylga ma'lumot yozadi
	file1.WriteString("this is file data2\n")
	defer file1.Close() //faylni yopadi
	// os.Remove("newfile.txt") //faylni o'chiradi
}

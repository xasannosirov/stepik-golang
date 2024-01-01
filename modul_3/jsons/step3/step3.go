package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func ThisMarshal() {
	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}
	// json.Marshal() structni json formatga o'tkazadi
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", data)
	// {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
}

func ThisMarshalIndent() {
	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}
	// json.MarshalIndent() struct jsonga chiroyli ko'rinishda o'tkazadi
	data, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)
	//{
	//	"Name": "John Connor",
	//	"Age": 35,
	//	"Status": true,
	//	"Values": [
	//		15,
	//		11,
	//		37
	//	]
	//}
}

func ThisUnmarshal() {
	data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	var s myStruct
	// json.Unmarshal() jsondan structga o'tkazadi
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v", s)
	// {John Connor 35 true [15 11 37]}
}

func ThisOther() {
	m := user{Name: "John Connor", Email: "test email"}

	data, _ := json.Marshal(m)
	
	//kesish uchun ishlatiladi 
	data = bytes.Trim(data, "{")
	fmt.Println(string(data))
	// json formatda yoki yo'qligini tekshiradi
	if !json.Valid(data) {
		fmt.Println("invalid json!")
	}

	fmt.Printf("%s", data)
	// вывод: "Name":"John Connor","Email":"test email","Status":false,"Language":null}
}

func main() {
	ThisMarshal()
	ThisMarshalIndent()
	ThisUnmarshal()
	ThisOther()
}

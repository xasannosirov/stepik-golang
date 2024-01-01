package main

import (
	"encoding/json"
	"fmt"
)

// json formatga o'girilganda qanday bo'lishini belgilab ketish
type myStruct struct {
	Name   string `json:"name"`
	Age    int    `json:",omitempty"`
	Status bool   `json:"-"`
}

func main() {
	m := myStruct{Name: "John Connor", Age: 0, Status: true}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // {"name":"John Connor"}
}

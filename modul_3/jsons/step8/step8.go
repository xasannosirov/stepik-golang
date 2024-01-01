package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// marshal va unmarshal o'rniga decode va encode dan foydalanish
func main() {
	var (
		src = testStruct{Name: "John Connor", Age: 35}
		dst = testStruct{}
		buf = new(bytes.Buffer)
	)

	enc := json.NewEncoder(buf)
	dec := json.NewDecoder(buf)

	enc.Encode(src)
	dec.Decode(&dst)

	fmt.Print(dst) // {John Connor 35}
}

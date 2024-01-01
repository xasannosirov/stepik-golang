package main

import "fmt"

func main() {
	// map e'lon qilishni 1-usuli
	var user = map[string]int{}
	user["one"] = 1
	fmt.Println(user)
	// map e'on qilishni 2-usuli
	user2 := make(map[string]int)
	user2["two"] = 2
	fmt.Println(user2)
	// map e'lon qilishni 3-usuli
	user3 := map[string]int{
		"three": 3,
	}
	fmt.Println(user3)
}

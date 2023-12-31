package main

import (
	"fmt"
	"strings"
)

// string lar bilan ishlashda strings package dagi funcksiyalardan foydalanish
func main() {
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"hello", "world"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("blanotblanot", "not", "***", -1),
		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
		strings.Trim("tetstet", "te"),
	)
}

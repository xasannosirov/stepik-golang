

package main

import (
	"fmt"
)

// fmt.Sprint() formatlab return qaytaradi
func main() {
	var a float64 = 100.123456789
	result := fmt.Sprintf("%.2f", a)
	fmt.Printf("%q", result) // "100.12"
}
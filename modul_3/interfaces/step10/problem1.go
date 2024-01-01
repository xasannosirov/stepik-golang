package main

import (
	"fmt"
	"os"
)

func printError(value interface{}) {
	fmt.Printf("value=%v: %T\n", value, value)
}

func readTask() (interface{}, interface{}, interface{}) {
	return 7.9, 8.9, "+"
}

func main() {
	value1, value2, operation := readTask()

	v1, ok := value1.(float64)
	if !ok {
		printError(value1)
		os.Exit(0)
	}

	v2, ok := value2.(float64)
	if !ok {
		printError(value2)
		os.Exit(0)
	}

	if v, ok := operation.(string); ok {
		switch v {
		case "+":
			fmt.Printf("%.4f\n", v1+v2)
		case "-":
			fmt.Printf("%.4f\n", v1-v2)
		case "*":
			fmt.Printf("%.4f\n", v1*v2)
		case "/":
			fmt.Printf("%.4f\n", v1/v2)
		default:
			fmt.Println("неизвестная операция")
			os.Exit(0)
		}
	} else {
		fmt.Println("неизвестная операция")
		os.Exit(0)
	}
}

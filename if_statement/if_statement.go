package main

import "fmt"

func main() {
	if_statement()
}

func if_statement() {
	// declare the variable and then check in the if statement
	a := 6
	b := 7
	if a < b {
		fmt.Println("a is less than b")
	}

	// declare the variable inside the if statement and then check
	if a, b := 7, 3; a > b {
		fmt.Println("a is greater than b")
	}

	// Checking multiple conditions in an if statement
	if a < b {
		fmt.Println("a is less than b")
	} else if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("a and b are equal")
	}

	// Checking multiple conditions in an if statement
	i := 4
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else if i == 4 {
		fmt.Println("Four")
	} else if i == 5 {
		fmt.Println("Five")
	}

	// checking cases using switch case
	i = 3
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

	// transition to the next state using the fallthrough keyword
	v := 42
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}

	// checking cases using switch case
	var c uint32
	fmt.Printf("Enter a number : ")
	fmt.Scan(&c)
	switch {
	case 1 <= c && c <= 9:
		fmt.Println("from 1 to 9")
	case 100 <= c && c <= 250:
		fmt.Println("from 100 to 250")
	case 1000 <= c && c <= 6000:
		fmt.Println("from 1000 to 6000")
	}
}
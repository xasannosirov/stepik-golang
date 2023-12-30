package main

import "fmt"

type Exmaple struct {
	On    bool
	Ammo  int
	Power int
}

func (c *Exmaple) Shoot() bool {
	if !c.On {
		return false
	} else if c.Ammo > 0 {
		c.Ammo -= 1
		return true
	}
	return false
}

func (c *Exmaple) RideBike() bool {
	if !c.On {
		return false
	} else if c.Power > 0 {
		c.Power -= 1
		return true
	}
	return false
}

func main() {
	ex := Exmaple{true, 1, 2}
	testStruct := &ex
	fmt.Println(testStruct)
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.RideBike())
}

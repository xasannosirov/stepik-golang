package main

import "fmt"

// struct yaratish
type Exmaple struct {
	On    bool
	Ammo  int
	Power int
}

// structga method yozish
func (c *Exmaple) Shoot() bool {
	if !c.On {
		return false
	} else if c.Ammo > 0 {
		c.Ammo -= 1
		return true
	}
	return false
}

// structga method yozish
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
	ex := Exmaple{true, 1, 2} //struct obyekti yasaldi
	testStruct := &ex //undan pointer olindi
	fmt.Println(*testStruct)
	fmt.Println(testStruct.Shoot()) //methodi ishga tushirildi
	fmt.Println(testStruct.Shoot()) //methodi ishga tushirildi
	fmt.Println(testStruct.RideBike()) //methodi ishga tushirildi 
	fmt.Println(testStruct.RideBike()) //methodi ishga tushirildi
	fmt.Println(testStruct.RideBike()) //methodi ishga tushirildi
}

package main

import "fmt"

type Teacher struct {
	LastName     string
	FirstName    string
	Age          int
	Students     int
	PhoneNumbers []string
}

func addTwoAge(t *Teacher) int {
	return t.Age + 2
} 

func (t *Teacher) ShowName() string {
	return t.FirstName + t.LastName
}

func (t *Teacher) ShowAge() int {
	return t.Age
}

func (t *Teacher) StudentsCount() int {
	return t.Students
}

func (t *Teacher) PhoneNumber() string {
	return t.PhoneNumbers[0]
}

func main() {
	tech := Teacher{
		FirstName:    "Xasan",
		LastName:     "Nosirov",
		Age:          17,
		Students:     22,
		PhoneNumbers: []string{
			"94-497-05-14", 
			"99-888-77-66",
		},
	}
	fmt.Println(tech)
	fmt.Println(tech.PhoneNumber())
	fmt.Println(tech.ShowAge())
	fmt.Println(tech.ShowName())
	fmt.Println(tech.StudentsCount())
	fmt.Println(addTwoAge(&tech))
}

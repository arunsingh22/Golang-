package main

import (
	"fmt"
)

type marks struct {
	english int
	hindi   int
	math    int
}

type student struct {
	Name  string
	Age   int
	Class string
	marks marks
}

func (s *student) String() string {
	return fmt.Sprintf("Name: %v, age: %v, passed/fail: %v", s.Name, s.Age, s.Class)
}

func main() {

	s1 := student{
		Name:  "arun",
		Age:   28,
		Class: "passed",
		marks: marks{
			english: 214,
			hindi:   12,
			math:    1241,
		},
	}

	fmt.Println(s1)

}

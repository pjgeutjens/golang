package main

import "fmt"

type myStruct struct {
	name string
	id   int
}

// structs are value types
// they are copied by value
// they are comparable
func main() {
	s := myStruct{
		name: "Pieter",
		id:   1,
	}
	s2 := s
	s.name = "Frederik"
	fmt.Println(s, s2)
}

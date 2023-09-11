package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func getUserName(u user) string {
	return u.name
}

func (u user) String() string {
	return u.name + " <" + u.email + ">"
}

func (u *user) SetName(name string) {
	u.name = name
}

func main() {
	u := user{name: "Pieter Jan Geutjens", email: "pieter.geutjens@gmail.com", age: 46}
	fmt.Println(getUserName(u))
	fmt.Println(u.String())
	u.SetName("Pieter Geutjens")
	fmt.Println(u.String())
}

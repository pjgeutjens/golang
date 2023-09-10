package main

import "fmt"

type printer interface {
	Print() string
}

type user struct {
	name  string
	email string
	age   int
}

type menuItem struct {
	name  string
	price float64
}

func (u user) Print() string {
	return u.name + " <" + u.email + ">"
}

func (m menuItem) Print() string {
	return fmt.Sprintf("%v - %.2f", m.name, m.price)
}

func main() {
	var p printer
	p = user{name: "Pieter Jan Geutjens", email: "pieter.geutjens@gmailcom", age: 46}
	fmt.Println(p.Print())
	p = menuItem{name: "Spaghetti", price: 12.50}
	fmt.Println(p.Print())

	// u := p.(user) // this will panic because p is not of type user
	u, ok := p.(user) // this will not panic because ok will be false
	fmt.Println(u, ok)

	mi, ok := p.(menuItem) // this will not panic because p is a menuItem
	fmt.Println(mi, ok)

	switch v := p.(type) {
	case user:
		fmt.Println("user:", v)
	case menuItem:
		fmt.Println("menuItem:", v)
	default:
		fmt.Println("unknown type:", v)
	}
}

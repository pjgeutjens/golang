package main

import (
	"errors"
	"fmt"
)

// error is an interface type that has an Error() method
// type error interface {
//   Error() string
// }

func main() {
	err := errors.New("this is an error")
	fmt.Println(err)

	err2 := fmt.Errorf("this error wraps the first one: %w", err)
	fmt.Println(err2)
}

type menu struct {
	items []string
}

// a function that returns an error if the input is invalid
func (m *menu) add(item string) error {
	if item == "" {
		return errors.New("item cannot be empty")
	}

	for _, i := range m.items {
		if item == i {
			return fmt.Errorf("item %s already exists", item)
		}
	}

	m.items = append(m.items, item)
	return nil
}

// a function that calls the add method deals with the popssible errors
func Add() {
	m := &menu{}
	err := m.add("burger")
	if err != nil {
		fmt.Println(err)
	}
}

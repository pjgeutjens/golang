package main

import "fmt"

func greet1(name1 string, name2 string) {
	fmt.Println(name1)
	fmt.Println(name2)
}

func greet2(name1, name2 string) {
	fmt.Println(name1)
	fmt.Println(name2)
}

func greet3(names ...string) {
	for _, n := range names {
		fmt.Println(n)
	}
}

func names(name string, otherName *string) {
	name = "New Name"
	*otherName = "Other new name"
}

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// be aware naked returns are considered to create potential confusion
func named_divide(a, b int) (result int, ok bool) {
	if b == 0 { // 0, false
		return
	}
	result = a / b
	ok = true
	return
	// optional: return a / b, true
}

func main() {
	name, otherName := "Pieter", "Frederik"
	fmt.Println(name)
	fmt.Println(otherName)
	names(name, &otherName)
	fmt.Println(name)
	fmt.Println(otherName)
	fmt.Println(add(4, 5))
	result, ok := divide(45, 5)
	fmt.Println(result, ok)
	fmt.Println(divide(40, 5))
	fmt.Println(divide(55, 0))
}

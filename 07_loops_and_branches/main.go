package main

import "fmt"

func main() {
	loops()
	ifelse()
	switches()
}

func switches() {
	i := 5
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2, 3:
		fmt.Println("i is 2 or 3")
	case 2 + 3, 2*i + 3:
		fmt.Println("i is 5 or 13")
	default:
		fmt.Println("i is not 1, 2, 3 or 5 or 13")
	}

	// logical switch
	switch i = 8; true {
	case i > 10:
		fmt.Println("i is greater than 10")
	case i > 5:
		fmt.Println("i is greater than 5")
	default:
		fmt.Println("i is less than 5")
	}

	// the true in the logical switch can be omitted
	switch i = 8; {
	// etc...
	}
}

func ifelse() {
	i := 5
	if i > 10 {
		fmt.Println("i is greater than 10")
	} else if i > 5 {
		fmt.Println("i is greater than 5")
	} else if i < 5 {
		fmt.Println("i is less than 5")
	} else {
		fmt.Println("i is 5")
	}

	// if with a short statement
	if j := 5; j < 10 {
		// j is only available in this scope
	} else if j > 5 {
		fmt.Println("j is greater than 5")
	}
}

func loops() {
	i := 1

	// infinite loop
	for {
		fmt.Println(i)
		i += 1
		if i > 10 {
			break
		}
	}

	// loop till condition
	i = 1
	for i < 5 {
		fmt.Println(i)
		i += 1
	}

	// counter based loop
	for i = 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// looping over a slice
	s := []string{"Coffee", "Tea", "Hot Chocolate"}
	for idx, v := range s {
		fmt.Println(idx, v)
	}

	// looping over a map
	m := map[string]int{"foo": 1, "bar": 2}
	for key, v := range m {
		fmt.Println(key, v)
	}
	// looping over a map with a blank identifier
	for _, v := range m {
		fmt.Println(v)
	}
}

package main

import (
	"fmt"

	"packages/mypkg"
)

// this is a single line comment
func main() {
	result := mypkg.Add(2, 4) // any statement can have a comment at the end
	fmt.Println(result)
	fmt.Println(mypkg.PI)
}

/*
this is a multi line comment
*/

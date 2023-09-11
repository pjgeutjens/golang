package main

import "fmt"

func main() {
	fmt.Println("Main 1")        // 1
	defer fmt.Println("Defer 1") // 4
	fmt.Println("Main 2")        // 2
	defer fmt.Println("Defer 2") // 3
}

// function main outputs this:
// Main 1
// Main 2
// Defer 2
// Defer 1

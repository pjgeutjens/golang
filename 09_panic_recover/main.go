package main

import "fmt"

func main() {
	_ = divide(10, 5)
	_ = safe_divide(10, 0) // this will panic due to divide by zero and recover
	_ = divide(10, 0)      // this will panic and will not recover
}

func divide(dividend, divisor int) int {
	return dividend / divisor
}

func safe_divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil { // check for normal execution
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}

package main

import "fmt"

func main() {
	_ = divide(10, 5)
	_ = divide_with_recover(10, 0) // this will panic due to divide by zero and recover
	_ = divide(10, 0)              // this will panic and will not recover

	_, err := divide_with_error(10, 0) // this will return an error and return to flow
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = divide_with_named_error(10, 0) // this will return an error and return to flow
	if err != nil {
		fmt.Println(err)
		return
	}
}

func divide(dividend, divisor int) int {
	return dividend / divisor
}

func divide_with_recover(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil { // check for normal execution
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}

func divide_with_error(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}

// this function allows us to influence the return value in a recovery scenario
func divide_with_named_error(dividend, divisor int) (result int, err error) {
	if divisor == 0 {
		err = fmt.Errorf("cannot divide by zero")
		result = 0
		return
	}
	return dividend / divisor, nil
}

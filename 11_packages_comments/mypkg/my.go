// package mypkg provides functionality specifically for meeeee
// it contains some math functions
package mypkg

// Add adds 2 ints and returns the result
func Add(a, b int) int {
	return a + b
}

// Divide divides two ints and returns the result and a bool to check
// for divide by zero
func Divide(dividend, divider int) (int, bool) {
	if divider == 0 {
		return 0, false
	}
	return dividend / divider, true
}

// Secret is a secret function that is not available outside this package
func secret() string {
	return "sst, I am  secret 'cause this function starts with a lowercase letter"
}

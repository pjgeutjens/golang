package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var (
	i    int
	x, y int     = 3, 4
	f    float64 = math.Sqrt(float64(x*x + y*y))
	zz   uint    = uint(f)
)

var (
	ToBe bool       = false
	MaXi uint64     = 1<<64 - 1
	z    complex128 = cmplx.Sqrt(-5 + 12i)
)

const PI = 3.14

func add(x, y int) int {
	return x + y
}

func sayHello() {
	fmt.Println("Hello")
}

func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	// see, no := here because x and y are considered to be declared at the top
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(1000))
	fmt.Printf("Now I have %v problems.\n", math.Sqrt(99))
	yy := "implicitly declared"
	fmt.Println(yy)
	fmt.Println(add(23, 44))
	sayHello()
	a, b := swap("Hello", "Golang")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

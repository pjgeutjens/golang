package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	// blank operator
	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this will not work (unused variable)
	// b, c, d, e, f := 0, 1, 2, 3, "happiness"
	// fmt.Println(b, c, d, f)

	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	// hex display
	i := 255
	fmt.Printf("Hex conv of '%d' is '%x'\n", i, i)

	// conversion
	y := 42
	z := 42.0

	fmt.Printf("%v is of type %T\n", y, y)
	fmt.Printf("%v is of type %T\n", z, z)

	var m float32 = 3.1415
	fmt.Printf("%v is of type %T\n", m, m)

	// this will not work as float32 is not compatible with float64
	// z = m
	// fmt.Printf("%v is of type %T\n", z, z)

	z = float64(m)
	fmt.Printf("%v is of type %T\n", z, z)
}

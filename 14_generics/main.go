package main

import "fmt"

func main() {
	testScores := []float64{98.9, 93.3, 77.1, 82, 83}
	moreTestScores := []float32{98.9, 93.3, 77.1, 82, 83}

	mapTestScores := map[string]float64{
		"Harry": 98.9,
		"Sally": 93.3,
		"Billy": 77.1,
		"Lisa":  82,
	}

	c := clone(testScores)
	fmt.Println(&testScores[0], &c[0], c)
	// c := clone(moreTestScores)   // this fails because clone is only for float64

	d := clone2(moreTestScores)
	e := clone2(testScores) // this works on both types because generics
	fmt.Println(&moreTestScores[0], &d[0], d, e)

	f := cloneMap(mapTestScores)
	fmt.Println(f)

	a1 := []int{1, 2, 3}
	a2 := []float64{1.1, 2.2, 3.3}
	a3 := []string{"one", "two", "three"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Println(s1, s2, s3)
}

// a clone function that is strongly linked to float64
func clone(s []float64) []float64 {
	c := make([]float64, len(s))
	copy(c, s)
	return c
}

// a generic clone function
func clone2[V any](s []V) []V {
	c := make([]V, len(s))
	copy(c, s)
	return c
}

// a generic clone function for maps, with comparable keys
func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

type addable interface {
	int | float64 | string
}

func add[V addable](items []V) V {
	var sum V
	for _, v := range items {
		sum += v
	}
	return sum
}

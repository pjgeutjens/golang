package main

import "fmt"

func main() {
	arrays()
	slices()
	maps()
}

// arrays
// array assignments are copy operations
// still can test eauqlity, under the hood this compares the
// data types, the length and the values of the arrays
func arrays() {
	var arr [3]int
	fmt.Println(arr)
	arr = [3]int{1, 2, 3}

	fmt.Println(arr[1])
	arr[1] = 99
	fmt.Println(arr)
	fmt.Println(len(arr))

	b := arr
	fmt.Println(b)
	fmt.Println(arr == b)
	arr[2] = 69
	fmt.Println(arr)
	fmt.Println(b)
	fmt.Println(arr == b)
}

// slices
// almost the same as arrays except no length specified at initialization
func slices() {
	s := []string{"Coffee", "Tea", "Hot Chocolate"}
	fmt.Println(s)

	fmt.Println(s[1])
	s[1] = "Chai Latte"
	fmt.Println(s)
	fmt.Println(len(s))
}

// maps
// maps are unordered key value pairs
// maps are reference types
// maps are not comparable
// maps that are assigned to each other point to the same underlying data structure
func maps() {
	var m map[string]int
	fmt.Println(m)
	m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m["foo"])
	m["bar"] = 99
	delete(m, "foo")
	m["baz"] = 420

	fmt.Println(m)
	fmt.Println(m["foo"]) // this is ok, always returns in this case 0
	v, ok := m["foo"]     // if it matters, this returns 0, false
	fmt.Println(v, ok)

	d := map[string][]string{
		"coffee": {"Coffee", "Espresso"},
		"tea":    {"Chai", "Chai Latte"},
	}
	fmt.Println(d)
	fmt.Println(d["tea"])

	d["other"] = []string{"Hot Chocolate"}
	fmt.Println(d)
	delete(d, "tea")
	fmt.Println(d)

	fmt.Println(d["tea"])
	w, ok := d["tea"]
	fmt.Println(w, ok)

	m2 := d
	m2["coffee"] = []string{"Coffee", "Espresso", "Latte"}
	d["tea"] = []string{"Chai", "Chai Latte", "Tea"}
	fmt.Println(d)
	fmt.Println(m2)
}

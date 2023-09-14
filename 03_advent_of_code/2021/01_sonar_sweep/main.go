package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	measurements, err := read_from_file("./hard.txt")
	if err != nil {
		panic(err)
	}

	prev, count := 999, 0

	for _, i := range measurements {
		if i > prev {
			count++
		}
		prev = i
	}
	fmt.Println(count)
}

func read_from_file(filepath string) ([]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

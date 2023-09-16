package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readings, _ := read_from_file("/home/pgs/personal/golang/03_advent_of_code/2021/03_binary_diagnostics/hard.txt")
	// convert string to binary
	fmt.Println(readings)
	total_number_of_readings := len(readings)
	reading_length := len(readings[0])

	var tally []int = make([]int, reading_length)
	for i := 0; i < total_number_of_readings; i++ {
		for j := 0; j < reading_length; j++ {
			if readings[i][j] == '1' {
				tally[j]++
			}
		}
	}
	fmt.Println(tally)
	gamma := ""
	epsilon := ""
	for i := range tally {
		if tally[i] > total_number_of_readings/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Print(g * e)
}

func read_from_file(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var readings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reading := scanner.Text()
		readings = append(readings, reading)
	}
	return readings, nil
}

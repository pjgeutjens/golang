package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readings, _ := read_from_file("/home/pgs/personal/golang/03_advent_of_code/2021/03_binary_diagnostics/simple.txt")
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

	fmt.Println(oxygen_generator_rating(readings))
}

func get_majority_bit(readings []string) ([]byte, error) {
	if len(readings) == 0 {
		return []byte{}, fmt.Errorf("empty readings")
	}

	reading_length := len(readings[0])
	var tally []int = make([]int, reading_length)
	for i := 0; i < len(readings); i++ {
		for j := 0; j < reading_length; j++ {
			if readings[i][j] == '1' {
				tally[j]++
			}
		}
	}
	result := make([]byte, reading_length)
	for i := range tally {
		if tally[i] < len(readings)/2 {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return result, nil
}

func oxygen_generator_rating(readings []string) string {
	reading_length := len(readings[0])
	for len(readings) > 1 {
		for j := 0; j < reading_length; j++ {
			for i := 0; i < len(readings); i++ {
				keep, err := get_majority_bit(readings)
				if err != nil {
					return ""
				}
				if readings[i][j] != keep[j] {
					readings[i] = readings[len(readings)-1]
					readings = readings[:len(readings)-1]
				}
			}
		}
	}
	return readings[0]
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

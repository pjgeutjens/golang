package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x   int
	y   int
	aim int
}
type Instruction struct {
	direction string
	amount    int
}

func (p *Position) Move(instruction Instruction) {
	switch instruction.direction {
	case "up":
		// p.y -= instruction.amount
		p.aim -= instruction.amount
	case "down":
		// p.y += instruction.amount
		p.aim += instruction.amount
	case "forward":
		p.x += instruction.amount
		p.y += p.aim * instruction.amount
	case "back":
		p.x += instruction.amount
	}
	fmt.Printf("moved %v by %d, now at x: %d, y, %d\n", instruction.direction, instruction.amount, p.x, p.y)
}

func (p Position) Report() {
	fmt.Println(p.x * p.y)
}

func main() {
	start := Position{0, 0, 0}
	instructions, err := read_input("./hard.txt")
	if err != nil {
		panic(err)
	}

	for _, i := range instructions {
		start.Move(i)
	}
	start.Report()
}

func read_input(filepath string) ([]Instruction, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(line[1])
		i := Instruction{direction: line[0], amount: a}
		result = append(result, i)
	}
	return result, nil
}

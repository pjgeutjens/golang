package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func (l Line) Print() {
	println(l.start.x, l.start.y, l.end.x, l.end.y)
}

func main() {
	lines, size, err := read_input("/home/pgs/personal/golang/03_advent_of_code/2021/05_hydrothermal_venture/simple.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(lines); i++ {
		lines[i].Print()
	}

	grid := make(Grid, size+1)
	for _, line := range lines {
		for _, point := range line.start.points_in_line(line.end) {
			grid.Mark(point)
		}
	}
	fmt.Println(grid.Count(2))
}

func (p Point) points_in_line(target Point) []Point {
	result := []Point{}
	if p.x == target.x && p.y == target.y {
		return result
	}
	if p.x == target.x {
		for i := p.y; i <= target.y; i++ {
			result = append(result, Point{p.x, i})
		}
	}
	if p.y == target.y {
		for i := p.x; i <= target.x; i++ {
			result = append(result, Point{i, p.y})
		}
	}
	return result
}

type Grid [][]int

func (g *Grid) Mark(p Point) {
	(*g)[p.x][p.y] += 1
}

func (g Grid) Count(min int) int {
	counter := 0
	for _, row := range g {
		for _, col := range row {
			if col >= min {
				counter++
			}
		}
	}
	return counter
}

func read_input(filepath string) ([]Line, int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	// var readings []string
	scanner := bufio.NewScanner(file)

	var result []Line
	max_coord := 0
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		start := strings.Split(points[0], ",")
		end := strings.Split(points[1], ",")
		start_x, _ := strconv.Atoi(start[0])
		if start_x > max_coord {
			max_coord = start_x
		}
		start_y, _ := strconv.Atoi(start[1])
		if start_y > max_coord {
			max_coord = start_y
		}
		end_x, _ := strconv.Atoi(end[0])
		if end_x > max_coord {
			max_coord = end_x
		}
		end_y, _ := strconv.Atoi(end[1])
		if end_y > max_coord {
			max_coord = end_y
		}
		result = append(result, Line{Point{start_x, start_y}, Point{end_x, end_y}})
	}
	return result, max_coord, nil
}

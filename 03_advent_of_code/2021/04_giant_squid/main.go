package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board [][]int

type Game struct {
	boards []Board
	draws  []int
}

func (g *Game) Mark(drawn int) int {
	sum := 0
	for _, b := range g.boards {
		sum += b.Mark(drawn)
	}
	return sum
}

func (b Board) Print() {
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
}

func (b *Board) Mark(drawn int) int {
	count := 0
	for i := 0; i < len(*b); i++ {
		for j := 0; j < len(*b); j++ {
			if (*b)[i][j] == drawn {
				(*b)[i][j] = -1
				count++
			}
		}
	}
	return count
}

func (b *Board) Score(drawn int) int {
	sum := 0
	for i := 0; i < len(*b); i++ {
		for j := 0; j < len(*b); j++ {
			if (*b)[i][j] != -1 {
				sum += (*b)[i][j]
			}
		}
	}
	return sum * drawn
}

func read_from_file(filepath string, g *Game) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// var readings []string
	scanner := bufio.NewScanner(file)
	// first line contains the draws
	err = read_draws_from_file(scanner, g)
	if err != nil {
		return err
	}
	err = read_boards_from_file(scanner, g)
	if err != nil {
		return err
	}
	return nil
}

func read_boards_from_file(scanner *bufio.Scanner, g *Game) error {
	for scanner.Scan() {
		read_board(g)
	}
	return nil
}

func read_board(g *Game) {
	panic("unimplemented")
}

func read_draws_from_file(scanner *bufio.Scanner, g *Game) error {
	scanner.Scan()
	drawstrings := strings.Split(scanner.Text(), ",")
	var draws []int
	for _, d := range drawstrings {
		v, err := strconv.Atoi(d)
		if err != nil {
			return err
		}
		draws = append(draws, v)
	}
	g.draws = draws
	return nil
}

func main() {
	g := &Game{}
	read_from_file("./simple.txt", g)
	fmt.Print(g.draws)
}

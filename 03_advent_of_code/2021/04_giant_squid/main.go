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
	wins   []bool
}

func (g Game) BoardsRemaining() int {
	counter := 1
	for _, w := range g.wins {
		if !w {
			counter++
		}
	}
	return counter
}

func (g *Game) Mark(drawn int) int {
	sum := 0
	for _, b := range g.boards {
		sum += b.Mark(drawn)
	}
	return sum
}

func (g *Game) Play() {
	done := false
	for _, d := range g.draws {
		if done {
			break
		}
		fmt.Println("Draw:", d)
		g.Mark(d)
		for i, b := range g.boards {
			if b.isBingo() && !g.wins[i] {
				b.Print()
				fmt.Printf("Bingo! on %d, %d boards remaining\n", i, g.BoardsRemaining())
				fmt.Println(b.Score(d))
				g.wins[i] = true
			}
			if g.BoardsRemaining() == 0 {
				done = true
			}
		}
	}
}

func (b Board) Print() {
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	fmt.Println()
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

func (b Board) isBingo() bool {
	return b.isRowBingo() || b.isColumnBingo()
}

func (b Board) isRowBingo() bool {
	for i := 0; i < len(b); i++ {
		if b.isRowFull(i) {
			return true
		}
	}
	return false
}

func (b Board) isColumnBingo() bool {
	for i := 0; i < len(b); i++ {
		if b.isColumnFull(i) {
			return true
		}
	}
	return false
}

func (b Board) isRowFull(row int) bool {
	for i := 0; i < len(b); i++ {
		if b[row][i] != -1 {
			return false
		}
	}
	return true
}

func (b Board) isColumnFull(col int) bool {
	for i := 0; i < len(b); i++ {
		if b[i][col] != -1 {
			return false
		}
	}
	return true
}

func (b *Board) AddRow(row []int) {
	*b = append(*b, row)
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
		read_board(scanner, g, 5)
	}
	return nil
}

func read_board(scanner *bufio.Scanner, g *Game, size int) {
	board := Board{}
	for i := 0; i < size; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		if len(line) == 0 {
			i--
			break
		}
		var row []int
		for _, v := range line {
			n, _ := strconv.Atoi(v)
			row = append(row, n)
		}
		board.AddRow(row)
	}
	g.boards = append(g.boards, board)
	g.wins = append(g.wins, false)
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
	read_from_file("./hard.txt", g)
	fmt.Print(g.draws)
	fmt.Println()
	for _, b := range g.boards {
		b.Print()
	}
	g.Play()
}

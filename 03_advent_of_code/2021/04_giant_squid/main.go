package main

import "fmt"

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

func main() {
	b := Board{}
	b.Print()
}

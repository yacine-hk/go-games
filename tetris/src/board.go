package main

import (
	"fmt"
)

const (
	COLUMNS = 11
	ROWS    = 10
)

type Board struct {
	cells    [][]string
	vertices []Vertex
}

func newBoard() *Board {
	cells := make([][]string, ROWS)
	for i := range cells {
		cells[i] = make([]string, COLUMNS)
	}

	return &Board{
		cells: cells,
	}
}

func (b *Board) update(locs []Vertex, sign string) {
	b.initialBoard()
	b.updateBoardLoc(locs, sign)
	b.drawBoard()
}

func (b *Board) initialBoard() {
	for i := 1; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			b.cells[i][j] = "-"
		}
	}

	for _, v := range b.vertices {
		b.cells[v.R][v.C] = "O"
	}
}

func (b *Board) updateBoardLoc(locs []Vertex, sign string) {
	for _, loc := range locs {
		b.cells[loc.R][loc.C] = sign
	}
}

func (b *Board) drawBoard() {
	for i := 1; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			fmt.Print(b.cells[i][j])
		}
		fmt.Print("\n")
	}

	rowLen := len(b.line())
	rowLine := b.line()
	if rowLen > 0 {
		locs := []Vertex{}
		for i := range rowLen {
			for j := 0; j < COLUMNS; j++ {
				loc := Vertex{R: rowLine[i], C: j}

				locs = append(locs, loc)
			}

		}
		for _, loc := range locs {
			for k, bVertex := range b.vertices {
				if bVertex == loc {
                    b.vertices[k] = loc
				}
			}
		}

		fmt.Println("LOCS: ", locs)
		fmt.Println("VBLOCS: ", b.vertices)
		b.update(locs, "-")
	}
}

func (b *Board) line() []int {
	lineRow := []int{}
	isLine := false
	for i := 1; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if b.cells[i][j] == "-" {
				isLine = false
				goto rowloop
			}
			isLine = true
		}
	rowloop:
		if isLine {
			lineRow = append(lineRow, i)
		}
	}
	return lineRow
}

package main

import "fmt"

type Board struct {
	cells [][]string
}

func newBoard(game *Game) *Board {
	rows := make([][]string, game.Height)
	for i := range rows {
		rows[i] = make([]string, game.Width)
	}
	return &Board{
		cells: rows,
	}
}

func (b *Board) init(game *Game) {
	for i := range game.Width {
		for j := range game.Height {
			b.cells[i][j] = ". "
		}
	}

}

func (b *Board) draw(game *Game) {
	for i := range game.Width {
		for j := range game.Height {
			fmt.Print(b.cells[i][j])
		}
		fmt.Print("\n")
	}

}
func (b *Board) update(game *Game) {
    for _, point := range game.Snake.Body[:len(game.Snake.Body) - 1] {
		b.cells[point.R][point.C] = "O "
	}
	b.cells[game.Food.R][game.Food.C] = "X "
}

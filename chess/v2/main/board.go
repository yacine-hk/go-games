package main

import (
	"fmt"
)

type Board struct {
	cells [][]string
}


func (b *Board) draw(dotCells [][]string, whitePlayer, blackPlayer Player) {
	for i := 0; i < 9; i++ {
		fmt.Print("---")
	}
	fmt.Print("\n")
	for i := 0; i < BOARDSIZE; i++ {
		fmt.Print(i, "  ")
	}
	fmt.Print("\n")
	for i := 0; i < BOARDSIZE; i++ {
		fmt.Print("===")
	}
	fmt.Print("==\n")
	for i := range dotCells {
		b.cells[i] = make([]string, len(dotCells[i]))
		copy(b.cells[i], dotCells[i])
	}

    for k, v := range whitePlayer.pieceLocation {
		b.cells[k.X][k.Y] = v.piece + whitePlayer.color
    }

    for k, v := range blackPlayer.pieceLocation {
		b.cells[k.X][k.Y] = v.piece + blackPlayer.color
    }

	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			fmt.Print(b.cells[i][j] + " ")
		}
		fmt.Print("|", i, "\n")
	}
	for i := 0; i < BOARDSIZE; i++ {
		fmt.Print("===")
	}
	fmt.Print("==\n")
}

func (b *Board) draw1(dotCells [][]string, whitePlayer, blackPlayer Player) {
	for i := 0; i < 9; i++ {
		fmt.Print("---")
	}
	fmt.Print("\n")
	for i := 0; i < BOARDSIZE; i++ {
		fmt.Print(i, "  ")
	}
	fmt.Print("\n")
	for i := 0; i < BOARDSIZE; i++ {
		fmt.Print("===")
	}
	fmt.Print("==\n")
	for i := range dotCells {
		b.cells[i] = make([]string, len(dotCells[i]))
		copy(b.cells[i], dotCells[i])
	}

	for _, k := range whitePlayer.pieces {
		b.cells[k.location.X][k.location.Y] = k.piece + whitePlayer.color
	}
	for _, k := range blackPlayer.pieces {
		b.cells[k.location.X][k.location.Y] = k.piece + blackPlayer.color
	}
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			fmt.Print(b.cells[i][j] + " ")
		}
		fmt.Print("|", i, "\n")
	}
	for i := 0; i < BOARDSIZE; i++ {
		fmt.Print("===")
	}
	fmt.Print("==\n")

    /*
	for k, v := range whitePlayer.pieceLocation {
		fmt.Println("WHITE => ", k, v)
	}

	for k, v := range blackPlayer.pieceLocation {
		fmt.Println("BLACK => ", k, v)
	}
    */
}

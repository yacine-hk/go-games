package main

import (
	"fmt"
)

const (
	WIDTH  = 3
	HEIGHT = 3
)

var natural = "\033[0;1;37;40m"
var yellow = "\033[0;1;33;40m"
var blue = "\033[0;1;34;40m"
var sign = "_"

type Board struct {
	cells [][]string
}

func newBoard() *Board {
	row := make([][]string, WIDTH)
	for i := range WIDTH {
		row[i] = make([]string, HEIGHT)
	}

	return &Board{
		cells: row,
	}
}

func (b *Board) clear() {
	fmt.Print("\033[H\033[2J")
}

func (b *Board) init(xPlayer, oPlayer Player) {
	for i := range WIDTH {
		for j := range HEIGHT {
			b.cells[i][j] = natural + sign + " "
		}
	}
	for k, v := range xPlayer.pieceLocs {
		b.cells[k.R][k.C] = blue + v
	}

	for k, v := range oPlayer.pieceLocs {
		b.cells[k.R][k.C] = yellow + v
	}
}

func (b *Board) draw(currentPlayer, xPlayer, oPlayer Player, cursorLoc Point) {
	for k, v := range xPlayer.pieceLocs {
		b.cells[k.R][k.C] = blue + v
	}

	for k, v := range oPlayer.pieceLocs {
		b.cells[k.R][k.C] = yellow + v
	}

	currentPlayer.ascii(b, cursorLoc, xPlayer, oPlayer)

	for i := range WIDTH {
		for j := range HEIGHT {
			fmt.Print(b.cells[i][j])
		}
		fmt.Print("\n")
	}

}

func (p *Player) ascii(b *Board, loc Point, xPlayer, oPlayer Player) {
    emptyPoint := Point{R:-1, C:-1}
	if loc == emptyPoint {
        return
	}
	highlight := p.piece.color
	cellSign := sign + " "
	v, ok := xPlayer.pieceLocs[loc]
	if ok {
		cellSign = v
	}
	v, ok = oPlayer.pieceLocs[loc]
	if ok {
		cellSign = v
	}
	b.cells[loc.R][loc.C] = highlight + cellSign + natural
}

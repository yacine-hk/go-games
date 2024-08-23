package main

import (
	"slices"
)

var colors = []string{"\033[0;1;32;44m", "\033[0;1;32;43m"}
var signs = []string{"X", "O"}

type Piece struct {
    loc Point
    color string
    sign string
}


func newPiece(sign string) *Piece {
    return &Piece {
        sign: sign,
        color: colors[slices.Index(signs, sign)],
    }
}

func setPiece(loc Point, sign string) *Piece {
    return &Piece {
        loc: loc,
        sign: sign,
        color: colors[slices.Index(signs, sign)],
    }
}

package main

import (
	"chess/moves"
	"fmt"
)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func remove(s []Piece, i int) []Piece {
     return append(s[:i], s[i+1:]...)
}

var dotCells = [][]string{
        {"--", "--", "--", "--", "--", "--", "--", "--"},
        {"--", "--", "--", "--", "--", "--", "--", "--"},
        {"--", "--", "--", "--", "--", "--", "--", "--"},
        {"--", "--", "--", "--", "--", "--", "--", "--"},
        {"--", "--", "--", "--", "--", "--", "--", "--"},
        {"--", "--", "--", "--", "--", "--", "--", "--"},
        {"--", "--", "--", "--", "--", "--", "--", "--"},
        {"--", "--", "--", "--", "--", "--", "--", "--"},
    }

var pieces = [6]string{"p", "r", "b", "n", "k", "q"}
var orderPieces = []string{"r", "n", "b", "q", "k", "b", "n", "r"}
const (
    boardSize = 8
    pawn = "p"
    rook = "r"
    bishop = "b"
    knight = "n"
    king = "k"
    queen = "q"
    white = "1"
    black = "2"
)


func piecesInitialLocs(player *Player, opponent Player) []Piece {
    var s, u = 0, 1
    if player.color == white {
        s = 7
        u = -1
    }
    locs := []Piece{}
    for i := 0; i < boardSize; i++ {
        locs = append(locs, Piece{
            piece: orderPieces[i], 
            location: moves.Vertex{X:s, Y:i}, 
            validMoves: generateAllMoves(orderPieces[i], moves.Vertex{X:s, Y:i}, opponent.color),
        })
        locs = append(locs, Piece{
            piece: "p", 
            location: moves.Vertex{X:s+u, Y:i}, 
            validMoves: generateAllMoves("p", moves.Vertex{X:s+u, Y:i}, opponent.color),
        })
    }
    return locs
}

var bblackSetLocs [][]moves.Vertex = [][]moves.Vertex{
    {{X:0, Y:3}, {X:4, Y:0}},
    {{X:0, Y:2}, {X:4, Y:4}},
    {{X:0, Y:1}, {X:4, Y:5}},
    {{X:1, Y:2}, {X:3, Y:2}},
    {{X:1, Y:3}, {X:3, Y:3}},
    {{X:1, Y:0}, {X:2, Y:0}},
    {{X:1, Y:7}, {X:0, Y:3}},
    {{X:1, Y:1}, {X:4, Y:1}},
}
var bwhiteSetLocs [][]moves.Vertex = [][]moves.Vertex{
    {{X:7, Y:3}, {X:5, Y:3}},
}


func main() {
    board := Board{}
    board.cells = make([][]string, len(dotCells))

    whitePlayer := Player{nameColor: "white", color:white}
    blackPlayer := Player{nameColor: "black", color:black}
    whitePlayer.pieces = piecesInitialLocs(&whitePlayer, blackPlayer)
    blackPlayer.pieces = piecesInitialLocs(&blackPlayer, whitePlayer)
    //blackPlayer.piecesSetLocs(bblackSetLocs)
    //whitePlayer.piecesSetLocs(bwhiteSetLocs)

    board.draw(dotCells, whitePlayer, blackPlayer)

    for {
        whitePlayer.rolePlayerMove(&blackPlayer, board)
        board.draw(dotCells, whitePlayer, blackPlayer)
        fmt.Println("mated?", blackPlayer.isCheckMated)
        if blackPlayer.isCheckMated == true {
            fmt.Println("OVER")
            return
        }

        blackPlayer.rolePlayerMove(&whitePlayer, board)
        board.draw(dotCells, whitePlayer, blackPlayer)
        fmt.Println("mated?", whitePlayer.isCheckMated)
        if whitePlayer.isCheckMated {
            fmt.Println("OVER")
            return
        }
    }

}

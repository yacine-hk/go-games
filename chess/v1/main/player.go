package main

import (
    "fmt"

    "chess/moves"
)

type Player struct {
    nameColor string
    color string
    pieces []Piece
    isCheckMated bool
}

func (p Player) Copy() Player {
    newPieces := make([]Piece, len(p.pieces))
    copy(newPieces, p.pieces)
    return Player {
        nameColor: p.nameColor,
        color: p.color,
        pieces: newPieces,
        isCheckMated: p.isCheckMated,
    }
}


func (p Player) rolePlayerMove(OpponentPlayer *Player, b Board) {
    var px, py, ux, uy int

    for {
        fmt.Printf("%s Move: ", p.nameColor)
        fmt.Scan(&px, &py, &ux, &uy)
        playerMove := moves.Vertex{X:ux, Y:uy}
        playerPieceLoc := moves.Vertex{X:px, Y:py}
        for i := range p.pieces {
            if p.pieces[i].location == playerPieceLoc {
                if b.verifyCheck(p, *OpponentPlayer, playerPieceLoc, playerMove) {
                    fmt.Println("OP CHECK")
                    b.draw(dotCells, p, *OpponentPlayer)
                    break
                }
                if !p.isValid(p.pieces[i].piece,playerPieceLoc, playerMove, *OpponentPlayer, b) {
                    fmt.Println("Invalid move")
                    break
                }
                p.pieces[i].location = playerMove
                for i := range OpponentPlayer.pieces {
                    if OpponentPlayer.pieces[i].location == playerMove {
                        OpponentPlayer.pieces = remove(OpponentPlayer.pieces, i)
                        return
                    }
                }
                if b.verifyCheck(*OpponentPlayer, p, playerPieceLoc, playerMove) {
                    fmt.Println("PLAY CHECK")
                    if OpponentPlayer.verifyCheckMate(p, b) {
                        OpponentPlayer.isCheckMated = true
                        fmt.Println("CHECKMATEE")
                        return
                    }

                }

                // Player takes the OpponentPlayer Piece
                
                return
            }
        }
        fmt.Println("Invalid Piece Location, try again")
    }
   
}


func (p Player) verifyCheckMate(checker Player, b Board) (bool) {
    //fmt.Println("VERIFYING")
    for _, k := range p.pieces {
        for _, j := range k.validMoves {
            //fmt.Println("P MOVES", k.piece, k.location, j, p.isValid(k.piece, k.location, j, checker, b), !b.verifyCheck(p, checker, k.location, j) )
            if p.isValid(k.piece, k.location, j, checker, b){
                if !b.verifyCheck(p, checker, k.location, j) {
                fmt.Println("NOT CHECKMATE MOVE", j, k.location)
                return false
                }
            }
        }
    }
    return true
}

func (p *Player) piecesSetLocs(pMove [][]moves.Vertex) {
    for i := range p.pieces {
        for _, j := range pMove {
            if p.pieces[i].location == j[0] {
                p.pieces[i].location = j[1]
            }
        }
    }
}

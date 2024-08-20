package moves

import "fmt"

type Vertex struct {
    X int
    Y int
}

const boardSize = 8

func isValid(move Vertex) (bool) {
    return move.X >= 0 && move.X < boardSize && move.Y >= 0 && move.Y < boardSize
}

func PawnMoves(loc Vertex, opponentColor string) ([]Vertex) {
    moves := []Vertex{}
    inc := 1
    line := 1
    if opponentColor == "2" {
        inc = -1
        line = 6
    }
    directions := []Vertex{
        {1, 0}, {2, 0}, {1, 1}, {1, -1},
    }
    for _, v := range directions {
        if loc.X != line && v.X == 2 {
            fmt.Println(loc)
            continue
        }
        move := Vertex{loc.X+(v.X*inc), loc.Y + (v.Y*inc)}
        if isValid(move) {
            moves = append(moves, move)
        }
    }
    return moves
}

func RookMoves(loc Vertex) ([]Vertex) {
    moves := []Vertex{}
    directions := []Vertex{
        {1, 0}, {-1, 0}, {0, 1}, {0, -1},
    }
    for _, v := range directions {
        for i := 1; i < boardSize; i++ {
            move := Vertex{loc.X + (v.X * i), loc.Y + (v.Y * i)}
            if isValid(move) {
                moves = append(moves, move)
            }
        }
    }
    return moves
}

func BishopMoves(loc Vertex) ([]Vertex) {
    moves := []Vertex{}
    directions := []Vertex{
        {1, 1}, {-1, 1}, {1, -1}, {-1, -1},
    }
    for _, v := range directions {
        for i := 1; i < boardSize; i++ {
            move := Vertex{loc.X + (v.X * i), loc.Y + (v.Y * i)}
            if isValid(move) {
                moves = append(moves, move)
            }
        }
    }
    return moves
}

func KnightMoves(loc Vertex) ([]Vertex) {
    moves := []Vertex{}
    directions := []Vertex{
        {2, 1}, {2, -1}, {-2, 1}, {-2, -1},
        {1, 2}, {1, -2}, {-1, 2}, {-1, -2},
    }
    for _, v := range directions {
        move := Vertex{loc.X + v.X, loc.Y + v.Y}
        if isValid(move) {
            moves = append(moves, move)
        }
    }
    return moves
}

func KingMoves(loc Vertex) ([]Vertex) {
    moves := []Vertex{}
    directions := []Vertex{
        {1, 0}, {-1, 0}, {0, 1}, {0, -1},
        {1, 1}, {-1, 1}, {1, -1}, {-1, -1},
    }

    for _, v := range directions {
        move := Vertex{loc.X + v.X, loc.Y + v.Y}
        if isValid(move) {
            moves = append(moves, move)
        }
    }
    return moves
}

func QueenMoves(loc Vertex) ([]Vertex) {
    moves := []Vertex{}
    moves = append(moves, BishopMoves(loc)...)
    moves = append(moves, RookMoves(loc)...)
    return moves
}


/*
old:
knight:

    if loc.X+2  >= 0 && loc.Y+1  >= 0 && loc.X+2 < boardSize  && loc.Y+1 < boardSize {
        moves = append(moves, Vertex{loc.X+2, loc.Y+1})
    }
    if loc.X+2  >= 0 && loc.Y+2  >= 0 && loc.X+2 < boardSize  && loc.Y+2 < boardSize {
        moves = append(moves, Vertex{loc.X+2, loc.Y+2})
    }
    if loc.X-2  >= 0 && loc.Y+1  >= 0 && loc.X-2 < boardSize  && loc.Y+1 < boardSize {
        moves = append(moves, Vertex{loc.X-2, loc.Y+1})
    }
    if loc.X-1  >= 0 && loc.Y+2  >= 0 && loc.X-1 < boardSize  && loc.Y+2 < boardSize {
        moves = append(moves, Vertex{loc.X-1, loc.Y+2})
    }
    if loc.X-2  >= 0 && loc.Y-1  >= 0 && loc.X-2 < boardSize  && loc.Y-1 < boardSize {
        moves = append(moves, Vertex{loc.X-2, loc.Y-1})
    }
    if loc.X-1  >= 0 && loc.Y-2  >= 0 && loc.X-1 < boardSize  && loc.Y-2 < boardSize {
        moves = append(moves, Vertex{loc.X-1, loc.Y-2})
    }
    if loc.X+2  >= 0 && loc.Y-1  >= 0 && loc.X+2 < boardSize  && loc.Y-1 < boardSize {
        moves = append(moves, Vertex{loc.X+2, loc.Y-1})
    }
    if loc.X+1  >= 0 && loc.Y-2  >= 0 && loc.X+1 < boardSize  && loc.Y-2 < boardSize {
        moves = append(moves, Vertex{loc.X+1, loc.Y-2})
    }

Vertex/ghlhg/
ertex Vertex


bishop
    for i := 0; i < boardSize - loc.X && i < boardSize - loc.Y; i++ {
        moves = append(moves, Vertex{loc.X+i, loc.Y+i})
    }
    for i := 0; i > boardSize - loc.X - boardSize - 1 && i > boardSize - loc.Y - boardSize - 1; i-- {
        moves = append(moves, Vertex{loc.X+i, loc.Y+i})
    }
    for i, j := 0, 0; i > boardSize - loc.X - boardSize - 1 && j < boardSize - loc.Y; i--{
        moves = append(moves, Vertex{loc.X+i, loc.Y+j})
        j++
    }
    for i, j := 0, 0; i > boardSize - loc.Y - boardSize - 1 && j < boardSize - loc.X; i--{
        moves = append(moves, Vertex{loc.X+j, loc.Y+i})
        j++
    }

rook

    for i := 0; i < boardSize - loc.X; i++ {
        moves = append(moves, Vertex{loc.X+i, loc.Y})
    }
    for i := 0; i > boardSize - loc.X - boardSize - 1; i-- {
        moves = append(moves, Vertex{loc.X+i, loc.Y})
    }
    for i := 0; i < boardSize - loc.Y; i++ {
        moves = append(moves, Vertex{loc.X, loc.Y+i})
    }
    for i := 0; i > boardSize - loc.Y - boardSize - 1; i-- {
        moves = append(moves, Vertex{loc.X, loc.Y+i})
    }

king

    if loc.X+1 >= 0 && loc.Y+1 >= 0 && loc.X+1 < boardSize  && loc.Y+1 < boardSize {
        moves = append(moves, Vertex{loc.X+1, loc.Y+1})
    }
    if loc.Y+1 >= 0 && loc.Y+1 < boardSize {
        moves = append(moves, Vertex{loc.X, loc.Y+1})
    }
    if loc.X+1 >= 0 && loc.X+1 < boardSize {
        moves = append(moves, Vertex{loc.X+1, loc.Y})
    }
    if loc.X-1 >= 0 && loc.Y+1 >= 0 && loc.X-1 < boardSize  && loc.Y+1 < boardSize {
        moves = append(moves, Vertex{loc.X-1, loc.Y+1})
    }
    if loc.X-1 >= 0 && loc.Y-1 >= 0 && loc.X-1 < boardSize  && loc.Y-1 < boardSize {
        moves = append(moves, Vertex{loc.X-1, loc.Y-1})
    }
    if loc.X+1 >= 0 && loc.Y-1 >= 0 && loc.X+1 < boardSize  && loc.Y-1 < boardSize {
        moves = append(moves, Vertex{loc.X+1, loc.Y-1})
    }
    if loc.X-1 >= 0 && loc.X-1 < boardSize {
        moves = append(moves, Vertex{loc.X-1, loc.Y})
    }
    if loc.Y-1 >= 0 && loc.Y-1 < boardSize {
        moves = append(moves, Vertex{loc.X, loc.Y-1})
    }

pawn
    moves = append(moves, Vertex{loc.X+(1*inc), loc.Y})
    moves = append(moves, Vertex{loc.X+(2*inc), loc.Y})
    moves = append(moves, Vertex{loc.X+(1*inc), loc.Y+(1*inc)})
*/


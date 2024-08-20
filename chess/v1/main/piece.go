package main

import "chess/moves"

type Piece struct {
    piece string
    location moves.Vertex
    validMoves []moves.Vertex
}



func generateAllMoves(piece string, loc moves.Vertex, opponentColor string) []moves.Vertex {
    pmoves := []moves.Vertex{}
    switch piece {
        case knight:
            pmoves = moves.KnightMoves(loc)
        case rook:
            pmoves = moves.RookMoves(loc)
        case bishop:
            pmoves = moves.BishopMoves(loc)
        case king:
            pmoves = moves.KingMoves(loc)
        case queen:
            pmoves = moves.QueenMoves(loc)
        case pawn:
            pmoves = moves.PawnMoves(loc, opponentColor)
    }
    return pmoves
    /*
    for _, k := range moves {
        fmt.Println(k.x, k.y)
    }
    */
}

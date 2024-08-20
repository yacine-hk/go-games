package main

type Piece struct {
	piece      string
	location   Vertex
	validMoves []Vertex
	isMoved    bool
	enPassant  bool
}

func generateAllMoves(piece string, loc Vertex, opponentColor string, opPlayer Player) []Vertex {
	pmoves := []Vertex{}
	switch piece {
	case KNIGHT:
		pmoves = KnightMoves(loc)
	case ROOK:
		pmoves = RookMoves(loc)
	case BISHOP:
		pmoves = BishopMoves(loc)
	case KING:
		pmoves = KingMoves(loc)
	case QUEEN:
		pmoves = QueenMoves(loc)
	case PAWN:
		pmoves = PawnMoves(loc, opponentColor, opPlayer)
	}
	return pmoves
}

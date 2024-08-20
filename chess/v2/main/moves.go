package main


type Vertex struct {
	X int
	Y int
}

func isNotOutOfRange(move Vertex) bool {
	return move.X >= 0 && move.X < BOARDSIZE && move.Y >= 0 && move.Y < BOARDSIZE
}

func PawnMoves(loc Vertex, opponentColor string, opPlayer Player) []Vertex {
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
			//fmt.Println(loc)
			continue
		}
		move := Vertex{loc.X + (v.X * inc), loc.Y + (v.Y * inc)}
		if isNotOutOfRange(move) {
			moves = append(moves, move)
		}
	}

	return moves
}

func RookMoves(loc Vertex) []Vertex {
	moves := []Vertex{}
	directions := []Vertex{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}
	for _, v := range directions {
		for i := 1; i < BOARDSIZE; i++ {
			move := Vertex{loc.X + (v.X * i), loc.Y + (v.Y * i)}
			if isNotOutOfRange(move) {
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func BishopMoves(loc Vertex) []Vertex {
	moves := []Vertex{}
	directions := []Vertex{
		{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
	}
	for _, v := range directions {
		for i := 1; i < BOARDSIZE; i++ {
			move := Vertex{loc.X + (v.X * i), loc.Y + (v.Y * i)}
			if isNotOutOfRange(move) {
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func KnightMoves(loc Vertex) []Vertex {
	moves := []Vertex{}
	directions := []Vertex{
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
	}
	for _, v := range directions {
		move := Vertex{loc.X + v.X, loc.Y + v.Y}
		if isNotOutOfRange(move) {
			moves = append(moves, move)
		}
	}
	return moves
}

func KingMoves(loc Vertex) []Vertex {
	moves := []Vertex{}
	directions := []Vertex{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
		{0, 2}, {0, -2},
	}

	for _, v := range directions {
		move := Vertex{loc.X + v.X, loc.Y + v.Y}
		if isNotOutOfRange(move) {
			moves = append(moves, move)
		}
	}
	return moves
}

func QueenMoves(loc Vertex) []Vertex {
	moves := []Vertex{}
	moves = append(moves, BishopMoves(loc)...)
	moves = append(moves, RookMoves(loc)...)
	return moves
}

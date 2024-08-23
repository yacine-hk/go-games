package main


var wins = [][]Point{
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{2, 0}, {1, 1}, {0, 2}},
}

type Player struct {
	sign      string
	piece     Piece
	pieceLocs map[Point]string
}

func newPlayer(sign string) *Player {
	moves := make(map[Point]string)
	return &Player{
		pieceLocs: moves,
		sign:      sign,
        piece: *newPiece(sign),
	}

}

func (p *Player) isValidMove(loc Point) bool {
	_, ok := p.pieceLocs[loc]
	return !ok
}

func (p *Player) roleMove(loc Point) {
	piece := setPiece(loc, p.sign)
	p.piece = *piece
	p.pieceLocs[loc] = p.sign + " "
}


func (p *Player) isWin() (bool, int) {
	scores := 0
	for i, winsPoint := range wins {
		for _, winPoint := range winsPoint {
			for piecePoint := range p.pieceLocs {
				if winPoint == piecePoint {
					scores++
                    break
				}
			}
		}
        if scores == 3 {
            return true, i
        }
        scores = 0
	}
	return false, -1
}

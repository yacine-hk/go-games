package main

var whiteEditorPieces = []Piece{
	{
        piece: "k",
        location: Vertex{X: 0, Y: 0},
    },
	{
        piece: "q",
        location: Vertex{X: 2, Y: 2},
    },
	{
        piece: "q",
        location: Vertex{X: 5, Y: 4},
    },
	{
        piece: "p",
        location: Vertex{X: 3, Y: 7},
    },
}

var blackEditorPieces = []Piece{
	{
        piece: "p",
        location: Vertex{X: 1, Y: 7},
    },
	{
        piece: "k",
        location: Vertex{X: 3, Y: 0},
    },
}

func piecesEditorLocs(player , opponent *Player, editorPieces []Piece) []Piece {
    for i, _ := range editorPieces {
        editorPieces[i].validMoves = generateAllMoves(editorPieces[i].piece, editorPieces[i].location, opponent.color, *opponent)
    }
    return editorPieces
}

func piecesInitialLocs(player *Player, opponent Player) []Piece {
	var s, u = 0, 1
	if player.color == WHITE {
		s = 7
		u = -1
	}
	locs := []Piece{}
	for i := 0; i < BOARDSIZE; i++ {
		locs = append(locs, Piece{
			piece:      orderPieces[i],
			location:   Vertex{X: s, Y: i},
			validMoves: generateAllMoves(orderPieces[i], Vertex{X: s, Y: i}, opponent.color, opponent),
		})
		locs = append(locs, Piece{
			piece:      PAWN,
			location:   Vertex{X: s + u, Y: i},
			validMoves: generateAllMoves(PAWN, Vertex{X: s + u, Y: i}, opponent.color, opponent),
		})
	}
	return locs
}

var bblackSetLocs [][]Vertex = [][]Vertex{
	{{X: 0, Y: 0}, {X: 2, Y: 7}},
}
var bwhiteSetLocs [][]Vertex = [][]Vertex{
	{{X: 6, Y: 4}, {X: 4, Y: 4}},
	{{X: 7, Y: 4}, {X: 3, Y: 0}},
}

var bblackSetLocs1 [][]Vertex = [][]Vertex{
	{{X: 0, Y: 3}, {X: 5, Y: 0}},
	{{X: 0, Y: 2}, {X: 0, Y: 3}},
	{{X: 0, Y: 1}, {X: 4, Y: 5}},
	{{X: 1, Y: 2}, {X: 3, Y: 2}},
	{{X: 1, Y: 3}, {X: 3, Y: 3}},
	{{X: 1, Y: 0}, {X: 3, Y: 0}},
	{{X: 1, Y: 7}, {X: 5, Y: 7}},
	{{X: 1, Y: 1}, {X: 4, Y: 1}},
}
var bwhiteSetLocs1 [][]Vertex = [][]Vertex{
	{{X: 7, Y: 3}, {X: 5, Y: 3}},
	{{X: 6, Y: 6}, {X: 5, Y: 6}},
}

func piecesSetLocs(pMove [][]Vertex, p, opPlayer *Player) {
	for i := range p.pieces {
		for _, j := range pMove {
			if p.pieces[i].location == j[0] {
				p.pieces[i].location = j[1]
				playerPiece := p.pieces[i]
				p.pieces[i].validMoves = generateAllMoves(playerPiece.piece, playerPiece.location, opPlayer.color, *opPlayer)
			}
		}
	}
}

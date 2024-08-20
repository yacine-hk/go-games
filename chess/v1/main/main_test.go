package main

import (
	//"fmt"
	"testing"

	"chess/moves"
)

//var whiteSetLocs [][]moves.Vertex = [][]moves.Vertex{[]moves.Vertex{moves.Vertex{0, 0}, moves.Vertex{5, 5}}}
var blackSetLocs = [][]moves.Vertex{
    {{X:1, Y:5}, {X:3, Y:5}},
}

var whiteSetLocs = [][]moves.Vertex{
    {{X:6, Y:4}, {X:4, Y:4}},
    {{X:7, Y:3}, {X:3, Y:7}},
}

/*
func (p *Player) piecesSetLocs(pMove [][]moves.Vertex) {
    for i := range p.pieces {
        for _, j := range pMove {
            if p.pieces[i].location == j[0] {
                p.pieces[i].location = j[1]
            }
        }
    }
}
    */


func TestVerifyCheck(t *testing.T) {
	// Define test cases
    testCases := []struct {
        name      string
        opponent  Player
        player    Player
        whiteSet [][]moves.Vertex
        blackSet [][]moves.Vertex
        pieceLoc  moves.Vertex
        moveLoc   moves.Vertex
        expected  bool
    }{
		{
			name: "not Check",
            whiteSet:  [][]moves.Vertex{
                {{X:6, Y:4}, {X:4, Y:4}},
                {{X:7, Y:3}, {X:3, Y:7}},
            },
            blackSet:  [][]moves.Vertex{
                {{X:1, Y:5}, {X:3, Y:5}},
                //{{X:1, Y:6}, {X:2, Y:6}},
            },
			pieceLoc: moves.Vertex{X: 3, Y: 7},
			moveLoc:  moves.Vertex{X: 0, Y: 4},
			expected: false,
		},
        
        /*
		{
			name: "No Check",
			pieceLoc: moves.Vertex{X: 7, Y: 7},
			moveLoc:  moves.Vertex{X: 0, Y: 4},
			expected: false,
		},
		{
			name: "Check",
            whiteSet:  [][]moves.Vertex{
                {{X:6, Y:4}, {X:4, Y:4}},
                {{X:7, Y:3}, {X:3, Y:7}},
            },
            blackSet:  [][]moves.Vertex{
                {{X:1, Y:5}, {X:2, Y:5}},
                {{X:1, Y:6}, {X:3, Y:6}},
            },
			pieceLoc: moves.Vertex{X: 3, Y: 7},
			moveLoc:  moves.Vertex{X: 0, Y: 4},
			expected: true,
		},
*/
	}



	// Run test cases
    whitePlayer := Player{nameColor: "white", color:white}
    blackPlayer := Player{nameColor: "black", color:black}

	for _, tc := range testCases {
        blackPlayer.pieces = piecesInitialLocs(&blackPlayer, whitePlayer)
        whitePlayer.pieces = piecesInitialLocs(&whitePlayer, blackPlayer)
        whitePlayer.piecesSetLocs(tc.whiteSet)
        blackPlayer.piecesSetLocs(tc.blackSet)

        tc.player = whitePlayer
        tc.opponent = blackPlayer

        /*
		t.Run(tc.name, func(t *testing.T) {
			board := &Board{}
            board.cells = make([][]string, len(dotCells))
            board.draw(dotCells, tc.player, tc.opponent)
			result := board.verifyCheck(tc.opponent, tc.player, tc.pieceLoc, tc.moveLoc)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})

        */
        t.Run(tc.name, func(t *testing.T) {
			board := &Board{}
            board.cells = make([][]string, len(dotCells))
            board.draw(dotCells, tc.player, tc.opponent)
			result := tc.opponent.verifyCheckMate(tc.player, *board)
            if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
        })
	}
}

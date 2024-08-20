package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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
	BOARDSIZE = 8
	PAWN      = "p"
	ROOK      = "r"
	BISHOP    = "b"
	KNIGHT    = "n"
	KING      = "k"
	QUEEN     = "q"
	WHITE     = "1"
	BLACK     = "2"
)

func main1() {
	board := Board{}
	board.cells = make([][]string, len(dotCells))

	whitePlayer := &Player{nameColor: "white", color: WHITE}
	blackPlayer := &Player{nameColor: "black", color: BLACK}

	whitePlayer.pieces = piecesInitialLocs(whitePlayer, *blackPlayer)
	blackPlayer.pieces = piecesInitialLocs(blackPlayer, *whitePlayer)

	//whitePlayer.pieces = piecesEditorLocs(whitePlayer, blackPlayer, whiteEditorPieces)
	//blackPlayer.pieces = piecesEditorLocs(blackPlayer, whitePlayer, blackEditorPieces)

	//piecesSetLocs(bblackSetLocs, blackPlayer, whitePlayer)
	//piecesSetLocs(bwhiteSetLocs, whitePlayer, blackPlayer)

	newMapPieceLoc(whitePlayer)
	newMapPieceLoc(blackPlayer)

	board.draw(dotCells, *whitePlayer, *blackPlayer)

	for {
		whitePlayer.rolePlayerMove(blackPlayer)
		board.draw(dotCells, *whitePlayer, *blackPlayer)
		if blackPlayer.isCheckMated {
			fmt.Println("OVER")
			return
		}

		blackPlayer.rolePlayerMove(whitePlayer)
		board.draw(dotCells, *whitePlayer, *blackPlayer)
		if whitePlayer.isCheckMated {
			fmt.Println("OVER")
			return
		}
	}

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func parseMoveToVertex(coordinate string) ([]Vertex, string, error) {
	var out []Vertex
	var err error
	promote := QUEEN
	coordinates := strings.Split(coordinate, " ")
	if len(coordinates) < 4 {
		return nil, "", fmt.Errorf("INVALID INPUT LENGTH (SHOULD BE 4 OR 5)")
	}
    

	move := make([]int, 4)
	for i, c := range coordinates[:4] {
		move[i], err = strconv.Atoi(c)
		if err != nil {
			return nil, "", fmt.Errorf("INVALID INPUT")
		}
	}

	out = append(out, Vertex{X: move[0], Y: move[1]})
	out = append(out, Vertex{X: move[2], Y: move[3]})

	if len(coordinates) == 5 {
		userPromote := string(coordinates[4])
		for _, piece := range pieces[1:] {
			if userPromote == piece {
				promote = userPromote
				break
			}
		}
	}

	return out, promote, nil
}

func main() {
	board := Board{}
	board.cells = make([][]string, len(dotCells))

	whitePlayer := &Player{nameColor: "white", color: WHITE}
	blackPlayer := &Player{nameColor: "black", color: BLACK}

	whitePlayer.pieces = piecesInitialLocs(whitePlayer, *blackPlayer)
	blackPlayer.pieces = piecesInitialLocs(blackPlayer, *whitePlayer)

	//whitePlayer.pieces = piecesEditorLocs(whitePlayer, blackPlayer, whiteEditorPieces)
	//blackPlayer.pieces = piecesEditorLocs(blackPlayer, whitePlayer, blackEditorPieces)

	//piecesSetLocs(bblackSetLocs, blackPlayer, whitePlayer)
	//piecesSetLocs(bwhiteSetLocs, whitePlayer, blackPlayer)
	newMapPieceLoc(whitePlayer)
	newMapPieceLoc(blackPlayer)


	clearScreen()
	board.draw(dotCells, *whitePlayer, *blackPlayer)

	for {
		whitePlayer.rolePlayerMove(blackPlayer)
		if blackPlayer.isCheckMated {
			fmt.Println("OVER")
			return
		}
		clearScreen()
		board.draw(dotCells, *whitePlayer, *blackPlayer)

		blackPlayer.rolePlayerMove(whitePlayer)
		if whitePlayer.isCheckMated {
			fmt.Println("OVER")
			return
		}
		clearScreen()
		board.draw(dotCells, *whitePlayer, *blackPlayer)

		time.Sleep(50 * time.Millisecond)
	}
}

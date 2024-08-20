package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Player struct {
	nameColor     string
	color         string
	pieces        []Piece
	isCheckMated  bool
	pieceLocation map[Vertex]*Piece
	//piecesLocation map[string][]Vertex
}

func newMapPieceLoc(player *Player) {
	player.pieceLocation = make(map[Vertex]*Piece)
	for _, p := range player.pieces {
		player.pieceLocation[p.location] = &p
	}
}

func (p Player) Copy() Player {
	newPieces := make([]Piece, len(p.pieces))
	copy(newPieces, p.pieces)
	newPieceLocMap := make(map[Vertex]*Piece, len(p.pieceLocation))
	for k, v := range p.pieceLocation {
		pieceCopy := *v
		newPieceLocMap[k] = &pieceCopy
	}
	return Player{
		nameColor:     p.nameColor,
		color:         p.color,
		pieces:        newPieces,
		isCheckMated:  p.isCheckMated,
		pieceLocation: newPieceLocMap,
	}
}

func (p *Player) isValid(pieceName string, loc, move Vertex, opponent *Player, promote string, showLogs bool) (bool, error) {
	//pmoves := []Vertex{}
	if loc == move {
		return false, fmt.Errorf("loc == move")

	}

	if showLogs {
		for k, v := range p.pieceLocation {
			fmt.Println("P: ", k, v)
		}

		for k, v := range opponent.pieceLocation {
			fmt.Println("OP: ", k, v)
		}
	}

	validPieceLoc := false
	piece, ok := p.pieceLocation[loc]
	if !ok {
		fmt.Println(loc)
		return false, fmt.Errorf("there is no piece here: %v", loc)
	}

	if showLogs {
		fmt.Println("RANGE MOVES: ", piece.validMoves)
		fmt.Println("MOVE: ", move)
	}

	for _, validMove := range piece.validMoves {
		if validMove == move {
			validPieceLoc = true
			break
		}
	}

	if !validPieceLoc {
		return false, fmt.Errorf("NOT VALID PIECE LOCATION")
	}

	_, ok = p.pieceLocation[move]
	if ok {
		return false, fmt.Errorf("SAME COLOR")
	}

	if pieceName == KNIGHT {
		return true, nil
	}

	xinc, yinc := 1, 1
	xrange := move.X - loc.X
	yrange := move.Y - loc.Y

	if xrange < 0 {
		xinc = -1
	}
	if yrange < 0 {
		yinc = -1
	}

	if showLogs {
		fmt.Println("RANGE: ", xrange, yrange)
	}
	i, j := 0, 0
	xrange = abs(xrange) - 1
	yrange = abs(yrange) - 1

	if showLogs {
		fmt.Println("RANGE: ", xrange, yrange)
	}
	for i < xrange || j < yrange {
		if i < xrange {
			i++
		}
		if j < yrange {
			j++
		}

		xloc := loc.X + (i * xinc)
		yloc := loc.Y + (j * yinc)
		if showLogs {
			fmt.Println("YOUR piece: ", xloc, yloc)
		}

		v, ok := p.pieceLocation[Vertex{X: xloc, Y: yloc}]
		if ok {
			return false, fmt.Errorf("THERES YOUR PIECE IN PATH: %v, %v, %v", v, xloc, yloc)
		}
		v, ok = opponent.pieceLocation[Vertex{X: xloc, Y: yloc}]
		//fmt.Println("OPPS", xloc, yloc, v)
		if ok {
			return false, fmt.Errorf("THERES OPPONENT PIECE IN PATH: %v, %v, %v", v, xloc, yloc)
		}
		if showLogs {
			fmt.Println("SPECS => ", pieceName, yrange)
		}
	}

	if pieceName == PAWN {
		pawnRow := 7
		if p.color == WHITE {
			pawnRow = 0
		}
		if move.X == pawnRow {
			v, _ := p.pieceLocation[loc]
			v.piece = promote
		}

		if yrange == -1 {
			v, ok := p.pieceLocation[move]
			if ok {
				return false, fmt.Errorf("YOU CAN'T TAKES YOUR PIECE: %v, %v", v, move)
			}
			v, ok = opponent.pieceLocation[move]
			if ok {
				return false, fmt.Errorf("YOU CAN'T TAKES FORWARD: %v, %v", v, move)
			}
		} else {
			pawnInc := -1
			if p.color == WHITE {
				pawnInc = 1
			}

			v, ok := opponent.pieceLocation[Vertex{X: move.X + pawnInc, Y: move.Y}]
			if ok {
				if v.enPassant {
					delete(opponent.pieceLocation, Vertex{X: move.X + pawnInc, Y: move.Y})
					return true, nil
				}
			}

			v, ok = opponent.pieceLocation[move]
			if !ok {
				return false, fmt.Errorf("NOTHING TO TAKES: %v, %v", v, move)
			}
		}
		for k, v := range opponent.pieceLocation {
			if v.enPassant == true {
				opponent.pieceLocation[k].enPassant = false
				fmt.Println("OP: ", opponent)
			}
		}

		if xrange == 1 {
			v, _ := p.pieceLocation[loc]
			v.enPassant = true
			return true, nil
		}
	}

	if pieceName == KING && yrange == 1 {
		if piece.isMoved {
			return false, fmt.Errorf("CAN'T CASTLE: king already moved")
		}
		yRookInc := 1
		yKingInc := -1
		if move.Y-loc.Y < 0 {
			yKingInc = 1
			yRookInc = -2
		}

		rookPiece, ok := p.pieceLocation[Vertex{X: move.X, Y: move.Y + yRookInc}]
		if rookPiece.isMoved {
			return false, fmt.Errorf("CAN'T CASTLE: rook already moved")
		}
		if !ok {
			return false, fmt.Errorf("CAN'T CASLTE: rook not found")
		}

		if verifyCheck(p, opponent, loc, loc) {
			return false, fmt.Errorf("CAN'T CASLTE: there is check for the king")
		}
		if verifyCheck(p, opponent, loc, move) {
			return false, fmt.Errorf("CAN'T CASLTE: there is check for the king")
		}
		if verifyCheck(p, opponent, loc, Vertex{X: move.X, Y: move.Y + yKingInc}) {
			return false, fmt.Errorf("CAN'T CASLTE: there is check for the king")
		}

		rookMove := Vertex{X: move.X, Y: move.Y + yKingInc}
		delete(p.pieceLocation, rookPiece.location)
		rookPiece.location = rookMove
		rookPiece.validMoves = generateAllMoves(rookPiece.piece, rookPiece.location, opponent.color, *opponent)
		rookPiece.isMoved = true
		p.pieceLocation[rookMove] = rookPiece
		piece.isMoved = true
	}

	return true, nil
}

func (p *Player) playTheMove(playerPieceLoc, playerMove Vertex, OpponentPlayer *Player, promote string) error {
	fmt.Println("P: ", playerPieceLoc, playerMove)

	pieceMap, ok := p.pieceLocation[playerPieceLoc]
	if !ok {
		return fmt.Errorf("THERE IS NO PIECE")
	}

	ok, err := p.isValid(pieceMap.piece, playerPieceLoc, playerMove, OpponentPlayer, promote, false)
	if !ok {
		return fmt.Errorf("CAN'T PLAY THE MOVE: %v", err)
	}

	if verifyCheck(p, OpponentPlayer, playerPieceLoc, playerMove) {
		return fmt.Errorf("OP CHECK")
	}

	// PLAY THE MOVE (SET THE PIECE LOCATION)
	v, _ := p.pieceLocation[playerPieceLoc]
	delete(p.pieceLocation, playerPieceLoc)
	pieceMap.location = playerMove
	pieceMap.validMoves = generateAllMoves(pieceMap.piece, playerMove, OpponentPlayer.color, *OpponentPlayer)
	pieceMap.isMoved = true
	pieceMap.enPassant = v.enPassant
	p.pieceLocation[playerMove] = pieceMap

	// TAKES THE PIECE
	pieceMap, ok = OpponentPlayer.pieceLocation[playerMove]
	if ok {
		delete(OpponentPlayer.pieceLocation, playerMove)
		return nil
	}
	return nil
}

func (p *Player) rolePlayerMove(OpponentPlayer *Player) {
	var px, py, ux, uy int

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s move: ", p.nameColor)
		move, _ := reader.ReadString('\n')
		move = strings.TrimSpace(move)

		pieceCoords, promote, err := parseMoveToVertex(move)
		if err != nil {
			fmt.Println("INPUT ERR: ", err)
			continue
		}

		px, py = pieceCoords[0].X, pieceCoords[0].Y
		ux, uy = pieceCoords[1].X, pieceCoords[1].Y
		playerMove := Vertex{X: ux, Y: uy}
		playerPieceLoc := Vertex{X: px, Y: py}

		err = p.playTheMove(playerPieceLoc, playerMove, OpponentPlayer, promote)
		if err != nil {
			fmt.Println("PLAY THE MOVE ERR: ", err)
			continue
		}

        /*
		if verifyStaleMate(OpponentPlayer, p) {
			OpponentPlayer.isCheckMated = true
			fmt.Println("STALEMATE")
			return
		}
        */

		if p.verifyCheckMate(OpponentPlayer, playerPieceLoc, playerMove) {
			OpponentPlayer.isCheckMated = true
			fmt.Println("CHECKMATEE", p.nameColor, "win")
			return
		}

		return
	}

}

func findPieceLoc(m map[Vertex]*Piece, piece string) (Vertex, bool) {
	for k, v := range m {
		if v.piece == piece {
			return k, true
		}
	}
	return Vertex{}, false
}

func verifyCheck(checked, checker *Player, pieceLoc, moveLoc Vertex) bool {
	newCheckedPlayer := checked
	newCheckerPlayer := checker

	// play the move of checked player
	pieceName := newCheckedPlayer.pieceLocation[pieceLoc]
	delete(newCheckedPlayer.pieceLocation, pieceLoc)
	newCheckedPlayer.pieceLocation[moveLoc] = pieceName

	// checked player can eat checker's piece
	var newCheckerPlayerPiece *Piece
	for k := range newCheckerPlayer.pieceLocation {
		if k == moveLoc {
			newCheckerPlayerPiece = newCheckerPlayer.pieceLocation[moveLoc]
			delete(newCheckerPlayer.pieceLocation, moveLoc)
			break
		}
	}

	kingLoc, ok := findPieceLoc(newCheckedPlayer.pieceLocation, KING)
	if !ok {
		panic("NO KING!")
	}

	for location, piece := range newCheckerPlayer.pieceLocation {
		ok, _ := newCheckerPlayer.isValid(piece.piece, location, kingLoc, newCheckedPlayer, QUEEN, false)
        fmt.Println("IN", kingLoc, location)
		if ok {
			// undo the checked player move
			pieceName = newCheckedPlayer.pieceLocation[moveLoc]
			delete(newCheckedPlayer.pieceLocation, moveLoc)
			newCheckedPlayer.pieceLocation[pieceLoc] = pieceName

			// recover the checker player piece
			if newCheckerPlayerPiece != nil {
				newCheckerPlayer.pieceLocation[moveLoc] = newCheckerPlayerPiece
			}

			return true
		}
	}

	// undo the checked player move
	pieceName = newCheckedPlayer.pieceLocation[moveLoc]
	delete(newCheckedPlayer.pieceLocation, moveLoc)
	newCheckedPlayer.pieceLocation[pieceLoc] = pieceName
	return false
}

func verifyCheck2(checked, checker *Player, pieceLoc, moveLoc Vertex) bool {
	newCheckedPlayer := checked.Copy()
	newCheckerPlayer := checker.Copy()

	// play the move of checked player
	pieceName := newCheckedPlayer.pieceLocation[pieceLoc]
	delete(newCheckedPlayer.pieceLocation, pieceLoc)
	newCheckedPlayer.pieceLocation[moveLoc] = pieceName

	// checked player can eat checker's piece
	var newCheckerPlayerPiece *Piece
	for k := range newCheckerPlayer.pieceLocation {
		if k == moveLoc {
			newCheckerPlayerPiece = newCheckerPlayer.pieceLocation[moveLoc]
			delete(newCheckerPlayer.pieceLocation, moveLoc)
			break
		}
	}

	kingLoc, ok := findPieceLoc(newCheckedPlayer.pieceLocation, KING)
	if !ok {
		panic("NO KING!")
	}

	for location, piece := range newCheckerPlayer.pieceLocation {
		ok, _ := newCheckerPlayer.isValid(piece.piece, location, kingLoc, &newCheckedPlayer, QUEEN, false)
		if ok {
			// undo the checked player move
			pieceName = newCheckedPlayer.pieceLocation[moveLoc]
			delete(newCheckedPlayer.pieceLocation, moveLoc)
			newCheckedPlayer.pieceLocation[pieceLoc] = pieceName

			// recover the checker player piece
			if newCheckerPlayerPiece != nil {
				newCheckerPlayer.pieceLocation[moveLoc] = newCheckerPlayerPiece
			}

			return true
		}
	}

	// undo the checked player move
	pieceName = newCheckedPlayer.pieceLocation[moveLoc]
	delete(newCheckedPlayer.pieceLocation, moveLoc)
	newCheckedPlayer.pieceLocation[pieceLoc] = pieceName
	return false
}

func verifyStaleMate(player, opponent *Player) bool {
	newOpponentPlayer := opponent.Copy()
	newPlayer := player.Copy()

	for loc, pieceMap := range newPlayer.pieceLocation {
		for _, move := range pieceMap.validMoves {
			err := newPlayer.playTheMove(loc, move, &newOpponentPlayer, QUEEN)
			if err != nil {
				fmt.Println("NOT STALE MATE", loc, move)
				time.Sleep(1 * time.Microsecond)
				return false
			}
		}
	}
	return true
}

func (p *Player) verifyCheckMate(checked *Player, loc, move Vertex) bool {
	newCheckedPlayer := checked.Copy()
	newCheckerPlayer := p.Copy()

	for loc, pieceMap := range newCheckedPlayer.pieceLocation {
		for _, move := range pieceMap.validMoves {
			err := newCheckedPlayer.playTheMove(loc, move, &newCheckerPlayer, QUEEN)
			if err == nil {
				return false
			}
		}
	}
	return true
}

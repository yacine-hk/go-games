package main

import (
	"chess/moves"
	"fmt"
	"strings"
)

type Board struct {
	cells [][]string
}

func (p *Player) isValid(piece string, loc, move moves.Vertex, opponent Player, b Board) bool {
	pmoves := []moves.Vertex{}
	if loc == move {
		fmt.Println("loc == move")
		return false
	}
	switch piece {
	case knight:
		pmoves = moves.KnightMoves(loc)
		for i := range pmoves {
			if pmoves[i].X == move.X && pmoves[i].Y == move.Y {
				if strings.Index(b.cells[move.X][move.Y], opponent.color) > -1 {
					return true
				}
				if strings.Index(b.cells[move.X][move.Y], "--") > -1 {
					return true
				}
				return false
			}
		}
		return false
	case rook:
		pmoves = moves.RookMoves(loc)
	case bishop:
		pmoves = moves.BishopMoves(loc)
	case king:
		pmoves = moves.KingMoves(loc)
	case queen:
		pmoves = moves.QueenMoves(loc)
	case pawn:
		pmoves = moves.PawnMoves(loc, opponent.color)
		//fmt.Println("MOVEE", loc, move, pmoves)
		for _, k := range pmoves {
			if k == move {
				if move.Y == loc.Y {
					fmt.Println("IN", move, loc, b.cells[move.X][move.Y])
					if strings.Index(b.cells[move.X][move.Y], opponent.color) > -1 {
						return false
					}
					return true
				}
				if move.Y != loc.Y {
					fmt.Println("NOT")
					if strings.Index(b.cells[move.X][move.Y], opponent.color) > -1 {
						return true
					}
					return false
				}
				return true
			}
		}
		/*
		   for _, k := range pmoves {
		       fmt.Println("outs", loc, k, opponent.color)

		       if k.Y == loc.Y {
		           if strings.Index(b.cells[move.X][move.Y], opponent.color) > -1 {
		               return false
		           }
		       }
		       if k.Y != loc.Y {
		       fmt.Println("NOT", k, loc)
		           if strings.Index(b.cells[move.X][move.Y], opponent.color) > -1 {
		               return true
		           }
		           return false
		       }
		   return true
		   }
		*/
	}
	for i := range pmoves {
		if pmoves[i] == move {

			if strings.Index(b.cells[move.X][move.Y], p.color) > -1 {
				fmt.Println("same color", move, p.color)
				fmt.Println("Same color")
				return false
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
			i, j := 0, 0
			xrange = abs(xrange) - 1
			yrange = abs(yrange) - 1
			for i < xrange || j < yrange {
				if i < xrange {
					i++
				}
				if j < yrange {
					j++
				}

				//fmt.Println(p.pieces[i].location, loc)
				for _, k := range p.pieces {
					if k.location.X == loc.X+(i*xinc) && k.location.Y == loc.Y+(j*yinc) {
						fmt.Println("Cannot play this move")
						return false
					}
				}
				for _, k := range opponent.pieces {
					if k.location.X == loc.X+(i*xinc) && k.location.Y == loc.Y+(j*yinc) {
						fmt.Println("Cannot play this move")
						return false
					}
				}
				/*
				   if b.cells[loc.X+(i*xinc)][loc.Y+(j*yinc)] != " ." {
				       fmt.Println("Cannot play this move")
				       return false
				   }
				*/
			}
			/*
			   if strings.Index(b.cells[loc.X+(i*xinc)][loc.Y+(j*yinc)], opponent.color) > -1 {
			       fmt.Println("TAKES")
			       return true
			   }
			*/

			return true
		}
	}
	return false
}

func (b *Board) verifyCheck(opponent, player Player, pieceLoc, moveLoc moves.Vertex) bool {
	newPlayer := opponent.Copy()
	newCheckerPlayer := player.Copy()
	for i := range newPlayer.pieces {
		if newPlayer.pieces[i].location == pieceLoc {
			newPlayer.pieces[i].location = moveLoc
			for i := len(newCheckerPlayer.pieces) - 1; i >= 0; i-- {
				if newCheckerPlayer.pieces[i].location == moveLoc && newPlayer.isValid(newPlayer.pieces[i].piece, newPlayer.pieces[i].location, moveLoc, newCheckerPlayer, *b) {
					newCheckerPlayer.pieces = remove(newCheckerPlayer.pieces, i)
					//fmt.Println("HO", newCheckerPlayer.pieces)
				}
			}

			//b.draw(dotCells, newPlayer, player)
			break
		}
	}
	var kingLoc moves.Vertex
	for _, k := range newPlayer.pieces {
		if k.piece == "k" {
			kingLoc = k.location
			break
		}
	}
	//fmt.Println(len(newCheckerPlayer.pieces))
	for _, k := range newCheckerPlayer.pieces {
		//fmt.Println("verify", k.piece, k.location, kingLoc, newPlayer)
		//fmt.Println("checker", newCheckerPlayer)
		if player.isValid(k.piece, k.location, kingLoc, newPlayer, *b) {
			return true
		}
	}
	return false
}

func (b *Board) draw(dotCells [][]string, whitePlayer, blackPlayer Player) {
	for i := 0; i < 9; i++ {
		fmt.Print("---")
	}
	fmt.Print("\n")
	for i := 0; i < boardSize; i++ {
		fmt.Print(i, "  ")
	}
	fmt.Print("\n")
	for i := 0; i < boardSize; i++ {
		fmt.Print("===")
	}
	fmt.Print("==\n")
	for i := range dotCells {
		b.cells[i] = make([]string, len(dotCells[i]))
		copy(b.cells[i], dotCells[i])
	}

	for _, k := range whitePlayer.pieces {
		b.cells[k.location.X][k.location.Y] = k.piece + whitePlayer.color
	}
	for _, k := range blackPlayer.pieces {
		b.cells[k.location.X][k.location.Y] = k.piece + blackPlayer.color
	}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			fmt.Print(b.cells[i][j] + " ")
		}
		fmt.Print("|", i, "\n")
	}
	for i := 0; i < boardSize; i++ {
		fmt.Print("===")
	}
	fmt.Print("==\n")
}

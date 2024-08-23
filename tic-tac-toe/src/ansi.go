package main

import (
	"fmt"
	"time"
)

var lines = []string{"-", "-", "-", "|", "|", "|", "\\", "/"}

func main() {
	var direction string
	oldState, err := enableRawMode()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer disableRawMode(oldState)

	board := newBoard()
	cursor := newCursor(Point{0, 0})
	xPlayer := newPlayer("X")
	oPlayer := newPlayer("O")
	players := []*Player{xPlayer, oPlayer}
	var currentPlayer *Player
	i := 0
	for {
		fmt.Println("I", i)
		currentPlayer = players[i%2]
		clearScreen()
		board.init(*xPlayer, *oPlayer)
		board.draw(*currentPlayer, *xPlayer, *oPlayer, cursor.loc)
		direction = handleInput()

		if direction == "CR" {
			currentLoc := cursor.loc
			if !xPlayer.isValidMove(currentLoc) || !oPlayer.isValidMove(currentLoc) {
				continue
			}
			currentPlayer.roleMove(currentLoc)
			if ok, winIdx := currentPlayer.isWin(); ok {
				for _, wv := range wins[winIdx] {
					currentPlayer.pieceLocs[wv] = lines[winIdx] + " "
				}
				clearScreen()
				board.init(*xPlayer, *oPlayer)
				board.draw(*currentPlayer, *xPlayer, *oPlayer, Point{-1,-1})
				fmt.Printf("%s Player win\n", currentPlayer.sign)
				return
			}
			i++
		}
		if i == 9 {
			clearScreen()
			board.init(*xPlayer, *oPlayer)
			board.draw(*currentPlayer, *xPlayer, *oPlayer, Point{-1,-1})
			fmt.Println("draw")
			return
		}

		cursor.moveCursor(direction)

		time.Sleep(0 * time.Second)
	}
}

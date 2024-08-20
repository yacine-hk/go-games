package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var sing = "O"

func main() {
	board := newBoard()
	block := newBlock("I", Vertex{R: 1, C: 5}, "UP")
	board.update(block.currentLoc, sing)

	for {
		fmt.Println("###############################")
		fmt.Printf("Board: %+v\n", board.vertices)
		if block.rowCollision(board) {
			fmt.Println("NEW")
			block = newBlock("I", Vertex{R: 1, C: 5}, "UP")
			board.update(block.currentLoc, sing)
			continue
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("move: ")
		move, _ := reader.ReadString('\n')
		move = strings.ToUpper(strings.TrimSpace(move))

		fmt.Println("THE MOVE: ", move)
		block = block.moveBlock(move, board)
		board.update(block.currentLoc, sing)
		fmt.Println("###############################")
		time.Sleep(1 * time.Millisecond)
	}
}

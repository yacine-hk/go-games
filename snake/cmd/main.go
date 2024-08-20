package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	// Enable raw mode for real-time input
	oldState, err := enableRawMode()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer disableRawMode(oldState)

	// Initial snake direction
	direction := "RIGHT"

	game := newGame(Point{2, 2})
	game.Food = newFood(game)
	board := newBoard(game)
	board.init(game)
	board.update(game)
	board.draw(game)

	// Game loop
    inc := 0
	for {
		clearScreen()
		fmt.Println("Snake is moving in direction:", direction)
        fmt.Println(inc)
        inc++
		game.Over = game.render(direction)
        if game.Over {
            fmt.Println("GAME OVER")
            return
        }

		// Check for user input (non-blocking)
		var buf [3]byte
		n, err := syscall.Read(int(os.Stdin.Fd()), buf[:])
		if err == nil && n > 0 {
            			// Handle WASD keys
			if n == 1 {
				switch buf[0] {
				case 'w':
					direction = "UP"
				case 's':
					direction = "DOWN"
				case 'a':
					direction = "LEFT"
				case 'd':
					direction = "RIGHT"
				}
			}

			// Handle arrow keys (multi-byte sequences)
			if n == 3 && buf[0] == 0x1b && buf[1] == '[' {
				switch buf[2] {
				case 'A': // Up arrow
					direction = "UP"
				case 'B': // Down arrow
					direction = "DOWN"
				case 'C': // Right arrow
					direction = "RIGHT"
				case 'D': // Left arrow
					direction = "LEFT"
				}
			}
		}
		board.init(game)
		board.update(game)
		board.draw(game)
		time.Sleep(200 * time.Millisecond)

	}
}

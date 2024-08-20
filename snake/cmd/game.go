package main


type Game struct {
	Snake  *Snake
	Food   Point
	Width  int
	Height int
	Over   bool
}

func newGame(snakePosition Point) *Game {
	return &Game{
		Snake:  newSnake(snakePosition),
		Width:  20,
		Height: 20,
		Over:   false,
	}

}

func snakeHeadDirection(game *Game, direction string) Point {
	headLocation := game.Snake.Body[0]
	switch direction {
	case "UP":
		headLocation.R -= 1
	case "LEFT":
		headLocation.C -= 1
	case "RIGHT":
		headLocation.C += 1
	case "DOWN":
		headLocation.R += 1
    default:
        newPoint := Point{}
        newPoint.R = headLocation.R - game.Snake.Body[1].R
        newPoint.C = headLocation.C - game.Snake.Body[1].C
        headLocation.R += newPoint.R
        headLocation.C += newPoint.C
	}
	return headLocation
}

func (g *Game) snakeCollision() bool {
    snakeHead := g.Snake.Body[0]
    if snakeHead.R >= g.Width || snakeHead.C >= g.Height {
        return true
    }

    if snakeHead.R < 0 || snakeHead.C < 0 {
        return true
    }

    for _, snakeBody := range g.Snake.Body[1:] {
        if snakeHead == snakeBody {
            return true
        }
    }

    return false
}

func (g *Game) render(direction string) bool {
	newDirection := snakeHeadDirection(g, direction)
	if newDirection == g.Snake.Body[1] && len(g.Snake.Body) > 2 {
        newDirection = snakeHeadDirection(g, "")
	}

	g.Snake.moveSnake(newDirection)
	if newDirection == g.Food {
		g.Food = newFood(g)
		g.Snake.growSnake()
	}
    return g.snakeCollision()
}

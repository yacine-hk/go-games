package main

import "math/rand"


func newFood(game *Game) Point {
	validPoints := []Point{}
	for i := 0; i < game.Width; i++ {
		for j := 0; j < game.Height; j++ {
			for _, point := range game.Snake.Body {
				if i == point.R && j == point.C {
					continue
				}
				validPoints = append(validPoints, Point{R: i, C: j})
			}
		}
	}

    randomIdx := rand.Intn(len(validPoints))
    return validPoints[randomIdx]
}

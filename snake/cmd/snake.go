package main

import "fmt"

type Point struct {
	R, C int
}

type Snake struct {
	Body      []Point
	Direction Point
	Grow      bool
}

func newSnake(position Point) *Snake {
	points := make([]Point, 2)
	points[0] = position
	points[1] = Point{position.R, position.C - 1}
	return &Snake{
		Body:      points,
		Direction: position,
		Grow:      false,
	}
}

func shiftBody(position Point, body []Point) []Point {
	headPosition := body[0]
	middlePosition := Point{}
	fmt.Println(body)
	for i, pos := range body {
		if i == 0 {
			continue
		}
		if i == 1 {
			middlePosition = pos
			body[i] = headPosition
			continue
		}
		body[i] = middlePosition
		middlePosition = pos
	}
	body[0] = position
	return body
}

func (s *Snake) growSnake() {
	s.Body = append(s.Body, Point{0, 0})
}

func (s *Snake) moveSnake(position Point) {
	s.Body = shiftBody(position, s.Body)
}

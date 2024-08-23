package main

type Point struct {
	R int
	C int
}

type Cursor struct {
	loc Point
}

func newCursor(loc Point) *Cursor {
	return &Cursor{
		loc: loc,
	}
}

func (c *Cursor) isValidLoc(loc Point) *Cursor {
	if loc.R >= HEIGHT || loc.R < 0 {
		return c
	}
	if loc.C >= WIDTH || loc.C < 0 {
		return c
	}
	c.loc = loc
	return c
}

func (c *Cursor) moveCursor(direction string) *Cursor {
	switch direction {
	case "UP":
		return c.isValidLoc(Point{R: c.loc.R - 1, C: c.loc.C})
	case "DOWN":
		return c.isValidLoc(Point{R: c.loc.R + 1, C: c.loc.C})
	case "LEFT":
		return c.isValidLoc(Point{R: c.loc.R, C: c.loc.C - 1})
	case "RIGHT":
		return c.isValidLoc(Point{R: c.loc.R, C: c.loc.C + 1})
	}
	return c
}

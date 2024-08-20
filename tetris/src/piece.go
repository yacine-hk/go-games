package main

import (
	"fmt"
)

func remove(s []Vertex, i int) []Vertex {
	return append(s[:i], s[i+1:]...)
}


var FACES = []string{"UP", "RIGHT", "BOTTOM", "LEFT"}

const (
	DOWN   = "D"
	LEFT   = "L"
	RIGHT  = "R"
	ROTATE = "RO"
)

func findIndex[T comparable](element T, slice []T) (int, error) {
	for i, el := range slice {
		if el == element {
			return i, nil
		}
	}

	return -1, fmt.Errorf("element not found")
}

type Vertex struct {
	R int
	C int
}

type Piece struct {
	name       string
	face       string
	currentLoc []Vertex
	rotations  [][]Vertex
}

func newBlock(name string, centerLoc Vertex, face string) *Piece {
	idx, err := findIndex(face, FACES)
	if err != nil {
		panic("invalid face")
	}

	var locs [][]Vertex
	switch name {
	case "T":
		locs = TBlockRotations(centerLoc)
	case "L":
		locs = LBlockRotations(centerLoc)
	case "J":
		locs = JBlockRotations(centerLoc)
	case "S":
		locs = SBlockRotations(centerLoc)
	case "Z":
		locs = ZBlockRotations(centerLoc)
	case "I":
		locs = IBlockRotations(centerLoc)
	case "O":
		locs = OBlockRotations(centerLoc)
	}

	currentLoc := locs[idx]

	rotations := [][]Vertex{}
	rotIdx := idx

	//fmt.Println("alllocs: ", locs)
	//fmt.Println("currentloc: ", currentLoc)
	for range len(locs) {
		rotations = append(rotations, locs[rotIdx])
		rotIdx++
		if rotIdx == len(locs) {
			rotIdx = 0
		}
	}

	return &Piece{
		face:       face,
		name:       name,
		currentLoc: currentLoc,
		rotations:  rotations,
	}
}

func (p *Piece) rotateBlock(board *Board) *Piece {
	faceIdx, _ := findIndex(p.face, FACES)

	//fmt.Println("ROT ALGO: ", len(p.rotations), p.face, faceIdx)
	if faceIdx+1 == len(p.rotations) {
		faceIdx = -1
	}

    newPiece := newBlock(p.name, p.currentLoc[0], FACES[faceIdx+1])
    if !AABB(newPiece.currentLoc, board) {
        return p
    }

	//fmt.Println("ROTS: ", p.rotations, FACES[faceIdx+1])
	return newPiece
}


func (p *Piece) moveBlock(direction string, board *Board) *Piece {
	var centerLoc Vertex

	switch direction {
	case DOWN:
		centerLoc = Vertex{
			R: p.currentLoc[0].R + 1, C: p.currentLoc[0].C,
		}

	case LEFT:
		centerLoc = Vertex{
			R: p.currentLoc[0].R, C: p.currentLoc[0].C - 1,
		}

	case RIGHT:
		centerLoc = Vertex{
			R: p.currentLoc[0].R, C: p.currentLoc[0].C + 1,
		}

	case ROTATE:
		return p.rotateBlock(board)

	default:
		return p
	}

	newPiece := newBlock(p.name, centerLoc, p.face)
	if !AABB(newPiece.currentLoc, board) {
		return p
	}

	return newPiece
}

func AABB(location []Vertex, board *Board) bool {
	maxColumn := location[0].C
	minColumn := location[0].C
	maxRow := location[0].R
	for _, vertex := range location {
		if vertex.C > maxColumn {
			maxColumn = vertex.C
		}
		if vertex.C < minColumn {
			minColumn = vertex.C
		}
		if vertex.R > maxRow {
			maxRow = vertex.R
		}

        for _, bVertex := range board.vertices {
            if vertex == bVertex {
                return false
            }
        }
	}

	if maxColumn >= COLUMNS || minColumn < 0 {
		return false
	}

	return true
}

func (p *Piece) rowCollision(board *Board) bool {
	maxRow := p.currentLoc[0].R

	for _, vertex := range p.currentLoc {
		if vertex.R > maxRow {
			maxRow = vertex.R
		}
	}

	if maxRow == ROWS-1 {
		board.vertices = append(board.vertices, p.currentLoc...)
		return true
	}

	for _, vertex := range board.vertices {
		for _, pVertex := range p.currentLoc {
			if pVertex.C == vertex.C {
				if pVertex.R == vertex.R-1 {
					board.vertices = append(board.vertices, p.currentLoc...)
					return true
				}
			}
		}
	}

	return false
}

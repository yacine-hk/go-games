package main


func TBlockRotations(centerLoc Vertex) [][]Vertex {
	rotations := [][]Vertex{
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R - 1, C: centerLoc.C},
		},
		{
			centerLoc,
			{R: centerLoc.R + 1, C: centerLoc.C},
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R, C: centerLoc.C + 1},
		},
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R + 1, C: centerLoc.C},
		},
		{
			centerLoc,
			{R: centerLoc.R + 1, C: centerLoc.C},
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R, C: centerLoc.C - 1},
		},
	}
	return rotations
}

func ZBlockRotations(centerLoc Vertex) [][]Vertex {
	rotations := [][]Vertex{
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R - 1, C: centerLoc.C - 1},
		},
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R - 1, C: centerLoc.C + 1},
			{R: centerLoc.R + 1, C: centerLoc.C},
		},
		{
			centerLoc,
			{R: centerLoc.R + 1, C: centerLoc.C},
			{R: centerLoc.R + 1, C: centerLoc.C + 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
		},
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R + 1, C: centerLoc.C - 1},
		},
	}

	return rotations
}

func SBlockRotations(centerLoc Vertex) [][]Vertex {
	rotations := [][]Vertex{
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R - 1, C: centerLoc.C - 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
		},
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R + 1, C: centerLoc.C + 1},
		},
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R + 1, C: centerLoc.C},
			{R: centerLoc.R + 1, C: centerLoc.C - 1},
		},
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C - 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R + 1, C: centerLoc.C},
		},
	}

	return rotations
}

func JBlockRotations(centerLoc Vertex) [][]Vertex {

	rotations := [][]Vertex{
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R - 1, C: centerLoc.C - 1},
		},
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R - 1, C: centerLoc.C + 1},
			{R: centerLoc.R + 1, C: centerLoc.C},
		},
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R + 1, C: centerLoc.C + 1},
		},
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R + 1, C: centerLoc.C},
			{R: centerLoc.R + 1, C: centerLoc.C - 1},
		},
	}

	return rotations
}

func LBlockRotations(centerLoc Vertex) [][]Vertex {

	rotations := [][]Vertex{
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R - 1, C: centerLoc.C + 1},
		},
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R + 1, C: centerLoc.C},
			{R: centerLoc.R + 1, C: centerLoc.C + 1},
		},
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R + 1, C: centerLoc.C - 1},
		},
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R - 1, C: centerLoc.C - 1},
			{R: centerLoc.R + 1, C: centerLoc.C},
		},
	}

	return rotations
}

func IBlockRotations(centerLoc Vertex) [][]Vertex {
	rotations := [][]Vertex{
		{
			centerLoc,
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R, C: centerLoc.C + 2},
			{R: centerLoc.R, C: centerLoc.C + 1},
		},
		{
            centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R + 1, C: centerLoc.C},
			{R: centerLoc.R + 2, C: centerLoc.C},
		},
        /*
		{
			{R: centerLoc.R, C: centerLoc.C},
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R, C: centerLoc.C + 2},
			{R: centerLoc.R, C: centerLoc.C + 1},
		},
		{
			{R: centerLoc.R, C: centerLoc.C - 1},
			{R: centerLoc.R - 1, C: centerLoc.C - 1},
			{R: centerLoc.R + 1, C: centerLoc.C - 1},
			{R: centerLoc.R + 2, C: centerLoc.C - 1},
		},
        */
	}

	return rotations
}

func OBlockRotations(centerLoc Vertex) [][]Vertex {
	rotations := [][]Vertex{
		{
			centerLoc,
			{R: centerLoc.R - 1, C: centerLoc.C},
			{R: centerLoc.R, C: centerLoc.C + 1},
			{R: centerLoc.R - 1, C: centerLoc.C + 1},
		},
	}

	return rotations

}

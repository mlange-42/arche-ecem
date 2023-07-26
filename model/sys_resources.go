package main

// GrassGrid holds entity IDs
type GrassGrid struct {
	Grid[float64]
}

// NewGrassGrid creates a new GrassGrid
func NewGrassGrid(width, height int) GrassGrid {
	return GrassGrid{
		NewGrid[float64](width, height),
	}
}

package main

import "github.com/mlange-42/arche/ecs"

// GrassGrid holds entity IDs
type GrassGrid struct {
	Grid[ecs.Entity]
}

// NewGrassGrid creates a new GrassGrid
func NewGrassGrid(width, height int) GrassGrid {
	return GrassGrid{
		NewGrid[ecs.Entity](width, height),
	}
}

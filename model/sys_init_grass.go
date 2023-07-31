package main

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitGrass system
type InitGrass struct{}

// Initialize the system
func (s *InitGrass) Initialize(w *ecs.World) {
	grid := ecs.GetResource[GrassGrid](w)

	generator := generic.NewMap2[Position, Vegetation](w)

	width, height := grid.Width(), grid.Height()
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			pos := Position{
				X: i,
				Y: j,
			}
			veg := Vegetation{
				Biomass: rand.Float64(),
			}
			generator.NewWith(&pos, &veg)
		}
	}
}

// Update the system
func (s *InitGrass) Update(w *ecs.World) {}

// Finalize the system
func (s *InitGrass) Finalize(w *ecs.World) {}

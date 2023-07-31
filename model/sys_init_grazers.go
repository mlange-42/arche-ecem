package main

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// InitGrazers system
type InitGrazers struct {
	Count int
}

// Initialize the system
func (s *InitGrazers) Initialize(w *ecs.World) {
	grid := ecs.GetResource[GrassGrid](w)

	generator := generic.NewMap2[Position, Energy](w)

	width, height := grid.Width(), grid.Height()
	len := s.Count
	for i := 0; i < len; i++ {
		pos := Position{
			X: rand.Intn(width),
			Y: rand.Intn(height),
		}
		energy := Energy{rand.Float64()*0.5 + 0.1}

		generator.NewWith(&pos, &energy)
	}
}

// Update the system
func (s *InitGrazers) Update(w *ecs.World) {}

// Finalize the system
func (s *InitGrazers) Finalize(w *ecs.World) {}

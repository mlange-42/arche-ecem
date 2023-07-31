package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// Metabolism system/submodel
type Metabolism struct {
	Rate   float64
	filter generic.Filter1[Energy]
}

// Initialize the system
func (s *Metabolism) Initialize(w *ecs.World) {
	s.filter = *generic.NewFilter1[Energy]()
}

// Update the system
func (s *Metabolism) Update(w *ecs.World) {
	query := s.filter.Query(w)

	for query.Next() {
		ene := query.Get()
		ene.Value -= s.Rate
	}
}

// Finalize the system
func (s *Metabolism) Finalize(w *ecs.World) {}

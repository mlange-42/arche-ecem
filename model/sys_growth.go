package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// Growth system/submodel
type Growth struct {
	Rate   float64
	filter generic.Filter1[Vegetation]
}

// Initialize the system
func (s *Growth) Initialize(w *ecs.World) {
	s.filter = *generic.NewFilter1[Vegetation]()
}

// Update the system
func (s *Growth) Update(w *ecs.World) {
	query := s.filter.Query(w)

	for query.Next() {
		veg := query.Get()
		veg.Biomass += s.Rate
		if veg.Biomass > 1 {
			veg.Biomass = 1
		}
	}
}

// Finalize the system
func (s *Growth) Finalize(w *ecs.World) {}

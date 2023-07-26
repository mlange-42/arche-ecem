package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// Mortality system/submodel
type Mortality struct {
	filter generic.Filter1[Energy]
	toDie  []ecs.Entity
}

// Initialize the system
func (s *Mortality) Initialize(w *ecs.World) {
	s.filter = *generic.NewFilter1[Energy]()
}

// Update the system
func (s *Mortality) Update(w *ecs.World) {
	query := s.filter.Query(w)

	for query.Next() {
		ene := query.Get()
		if ene.Value <= 0 {
			s.toDie = append(s.toDie, query.Entity())
		}
	}

	for _, e := range s.toDie {
		w.RemoveEntity(e)
	}
	s.toDie = s.toDie[:0]
}

// Finalize the system
func (s *Mortality) Finalize(w *ecs.World) {}

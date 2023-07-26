package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// PopulationObserver observer
type PopulationObserver struct {
	filter generic.Filter0
}

// Initialize observer
func (o *PopulationObserver) Initialize(w *ecs.World) {
	o.filter =
		*generic.NewFilter0().
			With(generic.T[Energy]())
}

// Update observer
func (o *PopulationObserver) Update(w *ecs.World) {}

// Header of the observer
func (o *PopulationObserver) Header() []string {
	return []string{"Population"}
}

// Values of the observer
func (o *PopulationObserver) Values(w *ecs.World) []float64 {
	query := o.filter.Query(w)
	pop := query.Count()
	query.Close()

	return []float64{float64(pop)}
}

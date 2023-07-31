package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// GrassBiomassObserver observer
type GrassBiomassObserver struct {
	filter generic.Filter1[Vegetation]
}

// Initialize observer
func (o *GrassBiomassObserver) Initialize(w *ecs.World) {
	o.filter = *generic.NewFilter1[Vegetation]()
}

// Update observer
func (o *GrassBiomassObserver) Update(w *ecs.World) {}

// Header of the observer (i.e. column names)
func (o *GrassBiomassObserver) Header() []string {
	return []string{"Biomass"}
}

// Values of the observer
func (o *GrassBiomassObserver) Values(w *ecs.World) []float64 {
	query := o.filter.Query(w)
	bm := 0.0

	for query.Next() {
		veg := query.Get()
		bm += veg.Biomass
	}

	return []float64{float64(bm)}
}

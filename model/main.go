package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche-pixel/plot"
	"github.com/mlange-42/arche-pixel/window"
	"github.com/mlange-42/arche/ecs"
)

func main() {
	m := model.New(ecs.NewConfig().WithCapacityIncrement(1024))
	m.TPS = 60

	m.Seed()

	grid := NewGrassGrid(100, 60)

	ecs.AddResource(&m.World, &grid)

	m.AddSystem(&InitGrazers{Count: 1000})

	m.AddSystem(&Metabolism{Rate: 0.01})
	m.AddSystem(&Mortality{})

	m.AddUISystem((&window.Window{
		Title:        "Population",
		Bounds:       window.B(1200, 660, 700, 340),
		DrawInterval: 25,
	}).With(&plot.TimeSeries{
		Observer: &PopulationObserver{},
	}))

	m.AddSystem(&system.FixedTermination{Steps: 10000})

	pixelgl.Run(m.Run)
}

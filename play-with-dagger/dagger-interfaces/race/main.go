package main

import (
	"context"
)

type Race struct{}

func (m *Race) Race(ctx context.Context) (Void, error) {
	return dag.Car().
		WithCar(dag.Audi().WithSpeed(100).WithModel("A4").WithColor("red").AsCarIcar()).
		WithCar(dag.Bmw().WithSpeed(120).WithModel("M5").WithColor("blue").AsCarIcar()).
		Race(ctx)
}

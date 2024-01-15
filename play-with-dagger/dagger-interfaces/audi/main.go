package main

import (
	"context"
	"fmt"
)

type Audi struct {
	Model string
	Speed int
	Color string
}

func (a *Audi) WithModel(model string) *Audi {
	a.Model = model

	return a
}

func (a *Audi) WithSpeed(speed int) *Audi {
	a.Speed = speed

	return a
}

func (a *Audi) WithColor(color string) *Audi {
	a.Color = color

	return a
}

func (a *Audi) Drive(ctx context.Context) error {
	fmt.Printf(
		"A %s Audi model %s is driving at %dkmh..",
		a.Color,
		a.Model,
		a.Speed,
	)

	return nil
}

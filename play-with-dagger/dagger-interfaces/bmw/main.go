package main

import (
	"context"
	"fmt"
)

type Bmw struct {
	Model string
	Speed int
	Color string
}

func (a *Bmw) WithModel(model string) *Bmw {
	a.Model = model

	return a
}

func (a *Bmw) WithSpeed(speed int) *Bmw {
	a.Speed = speed

	return a
}

func (a *Bmw) WithColor(color string) *Bmw {
	a.Color = color

	return a
}

func (a *Bmw) Drive(ctx context.Context) error {
	fmt.Printf(
		"A %s Bmw model %s is driving at %dkmh..",
		a.Color,
		a.Model,
		a.Speed,
	)

	return nil
}


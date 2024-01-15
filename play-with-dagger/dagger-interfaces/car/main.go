package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Icar interface {
	DaggerObject

	WithModel(model string) Icar
	WithSpeed(speed int) Icar
	WithColor(color string) Icar
	Drive(ctx context.Context) error
}

type Car struct {
	Cars []Icar
}

func (c *Car) WithCar(car Icar) *Car {
	c.Cars = append(c.Cars, car)

	return c
}

func (c *Car) Race(ctx context.Context) error {
	g, gctx := errgroup.WithContext(ctx)


	for _, car := range c.Cars {
		car := car // Create a local variable and assign the value of car to it

		g.Go(func() error {
			if err := car.Drive(gctx); err != nil {
				return err
			}

			return nil
		})
	}

	return g.Wait()
}
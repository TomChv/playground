package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type ICar interface {
	DaggerObject

	WithModel(model string) ICar
	WithSpeed(speed int) ICar
	WithColor(color string) ICar
	Drive(ctx context.Context) error
}

type Car struct {
	Cars []ICar
}

func (c *Car) WithCar(car ICar) *Car {
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
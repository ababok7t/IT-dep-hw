package shapes

import (
	"errors"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (c Circle) GetArea() (float64, error) {
	if c.Radius <= 0 {
		err := errors.New("circle doesn't exist")
		return 0, err
	}
	area := math.Pi * c.Radius * c.Radius
	return area, nil
}

func (r Rectangle) GetArea() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		err := errors.New("rectangle doesn't exist")
		return 0, err
	}
	area := r.Width * r.Height
	return area, nil
}

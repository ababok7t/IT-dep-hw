package calculator

import (
	"errors"
	"hw2/shapes"
)

func TotalArea(circles []shapes.Circle, rectangles []shapes.Rectangle) (float64, error) {
	var sum float64
	sum = 0
	if len(circles) == 0 && len(rectangles) == 0 {
		err := errors.New("no shapes in file")
		return 0, err
	}
	for i := range circles {
		area, err := circles[i].GetArea()
		if err != nil {
			return 0, err
		}
		sum += area
	}
	for i := range rectangles {
		area, err := rectangles[i].GetArea()
		if err != nil {
			return 0, err
		}
		sum += area
	}
	return sum, nil
}

package calculator

import (
	"hw2/shapes"
	"testing"
)

func TestCalculator1(t *testing.T) {
	circles := []shapes.Circle{{Radius: 3.}, {Radius: 5.}, {Radius: 4.64}}
	rectangles := []shapes.Rectangle{{Width: 3, Height: 8}, {Width: 5, Height: 5}, {Width: 7.2, Height: 9.8}}
	expected := 294.0113834167798
	actual, err := TotalArea(circles, rectangles)

	if err != nil {
		t.Error(err.Error())
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

func TestCalculator2(t *testing.T) {
	circles := []shapes.Circle{{Radius: 3.}, {Radius: -5.}, {Radius: 4.64}}
	rectangles := []shapes.Rectangle{{Width: 3, Height: 0}, {Width: -3.123, Height: 5}, {Width: -9.875, Height: 9.8}}
	expected := 0.
	actual, err := TotalArea(circles, rectangles)

	if err == nil {
		t.Error("expected error")
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

func TestCalculator3(t *testing.T) {
	var circles []shapes.Circle
	var rectangles []shapes.Rectangle
	expected := 0.
	actual, err := TotalArea(circles, rectangles)

	if err == nil {
		t.Error("expected error")
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

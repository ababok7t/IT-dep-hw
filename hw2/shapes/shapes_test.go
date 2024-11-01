package shapes

import (
	"math"
	"testing"
)

func TestCircle1(t *testing.T) {
	radius := 17.
	expected := math.Pi * radius * radius
	circle := Circle{Radius: radius}
	actual, err := circle.GetArea()

	if err != nil {
		t.Error(err.Error())
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

func TestCircle2(t *testing.T) {
	radius := -5.
	expected := 0.
	circle := Circle{Radius: radius}
	actual, err := circle.GetArea()

	if err == nil {
		t.Error("expected error")
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

func TestCircle3(t *testing.T) {
	radius := 0.
	expected := 0.
	circle := Circle{Radius: radius}
	actual, err := circle.GetArea()

	if err == nil {
		t.Error("expected error")
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

func TestRectangle1(t *testing.T) {
	width := 3.
	height := 4.
	expected := width * height
	rectangle := Rectangle{Width: width, Height: height}
	actual, err := rectangle.GetArea()

	if err != nil {
		t.Error("error: " + err.Error())
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

func TestRectangle2(t *testing.T) {
	width := 7.
	height := -8.
	expected := 0.
	rectangle := Rectangle{Width: width, Height: height}
	actual, err := rectangle.GetArea()

	if err == nil {
		t.Error("expected error")
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

func TestRectangle3(t *testing.T) {
	width := -9.
	height := 2.
	expected := 0.
	rectangle := Rectangle{Width: width, Height: height}
	actual, err := rectangle.GetArea()

	if err == nil {
		t.Error("expected error")
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

func TestRectangle4(t *testing.T) {
	width := 0.
	height := 0.
	expected := 0.
	rectangle := Rectangle{Width: width, Height: height}
	actual, err := rectangle.GetArea()

	if err == nil {
		t.Error("expected error")
	}

	if actual != expected {
		t.Error("incorrect result")
	}
}

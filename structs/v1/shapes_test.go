package main

import (
	"testing"
)

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}
func Area(width float64, height float64) float64 {
	return width * height
}
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func testArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

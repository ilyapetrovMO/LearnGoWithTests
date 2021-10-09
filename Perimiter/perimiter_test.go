package main

import "testing"

var rectangle = &Rectangle{Width: 12.0, Height: 6.0}
var circle = &Circle{10}

func TestPerimiter(t *testing.T) {
	got := rectangle.Perimiter()
	want := 36.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Height: 12, Width: 6}, want: 36.0},
	}

	for _, tc := range areaTests {
		t.Run(tc.name, func(t *testing.T) {

			got := tc.shape.Area()

			if got != tc.want {
				t.Errorf("%#v got %.2f want %.2f", tc.shape, got, tc.want)
			}
		})
	}
}

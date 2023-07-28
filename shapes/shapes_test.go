package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.expected {
			t.Errorf("got %g want %g", got, tt.expected)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		name    string
		hasArea float64
	}{
		{Rectangle{12.0, 6.0}, "Rectangle", 72.0},
		{Circle{10}, "Circle", 314.1592653589793},
		{Triangle{12, 6}, "Triangle", 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.name, got, tt.hasArea)
			}
		})
	}
}

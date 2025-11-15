package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		name      string
		rectangle Rectangle
		expected  float64
	}{
		{
			name: "calculating a positive Perimeter",
			rectangle: Rectangle{
				Width:  10.0,
				Height: 10.0,
			},
			expected: 40.0,
		},
		{
			name: "calculating a zero value Perimeter",
			rectangle: Rectangle{
				Width:  0.0,
				Height: 0.0,
			},
			expected: 0.0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Perimeter(tc.rectangle)

			if result != tc.expected {
				t.Errorf("Perimeter(%v) = %v; want %v", tc.rectangle, result, tc.expected)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

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

	testCases := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tc := range testCases {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.expected {
				t.Errorf("%#v got %g want %g", tc.shape, got, tc.expected)
			}
		})

	}

}

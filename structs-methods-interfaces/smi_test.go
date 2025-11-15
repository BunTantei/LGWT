package structsmethodsinterfaces

import "testing"

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
		name      string
		rectangle Rectangle
		expected  float64
	}{
		{
			name: "calculating a positive Area",
			rectangle: Rectangle{
				Width:  12.0,
				Height: 6.0,
			},
			expected: 72.0,
		},
		{
			name: "calculating a zero value Area",
			rectangle: Rectangle{
				Width:  0.0,
				Height: 0.0,
			},
			expected: 0.0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Area(tc.rectangle)

			if result != tc.expected {
				t.Errorf("Area(%g) = %g; want %g", tc.rectangle, result, tc.expected)
			}
		})
	}
}

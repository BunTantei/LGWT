package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{
			name:     "calculating a positive Perimeter",
			a:        10.0,
			b:        10.0,
			expected: 40.0,
		},
		{
			name:     "calculating a zero value Perimeter",
			a:        0.0,
			b:        0.0,
			expected: 0.0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Perimeter(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("Perimeter(%.2f, %.2f) = %.2f; want %.2f", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{
			name:     "calculating a positive Area",
			a:        12.0,
			b:        6.0,
			expected: 72.0,
		},
		{
			name:     "calculating a zero value Area",
			a:        0.0,
			b:        0.0,
			expected: 0.0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Area(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("Area(%.2f, %.2f) = %.2f; want %.2f", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

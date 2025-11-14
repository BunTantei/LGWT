package structsmethodsinterfaces

import "testing"

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimerter(t *testing.T) {
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
		{
			name:     "calculating a negative value Perimeter",
			a:        -10.0,
			b:        -10.0,
			expected: -40.0,
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

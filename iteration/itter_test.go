package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	testCases := []struct {
		name     string
		a        string
		expected string
	}{
		{
			name:     "5 a",
			a:        "a",
			expected: "aaaaa",
		},
		{
			name:     "5 1",
			a:        "1",
			expected: "11111",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Repeat(tc.a)

			if result != tc.expected {
				t.Errorf("Repeat(%s) = %s; want %s", tc.a, result, tc.expected)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

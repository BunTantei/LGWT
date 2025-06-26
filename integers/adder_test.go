package integers

import "testing"

func TestAdder(t *testing.T) {
	// Define the "table" of test cases.
	// Each element in this slice is a struct representing one test case.
	testCases := []struct {
		name     string // A descriptive name for the test case
		a        int    // Input 'a' for the Add function
		b        int    // Input 'b' for the Add function
		expected int    // The expected result from Add(a, b)
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "negative numbers",
			a:        -2,
			b:        -3,
			expected: -5,
		},
		{
			name:     "positive and negative",
			a:        5,
			b:        -2,
			expected: 3,
		},
		{
			name:     "adding zero",
			a:        10,
			b:        0,
			expected: 10,
		},
		{
			name:     "zero plus zero",
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	// Iterate over each test case in our table.
	for _, tc := range testCases {
		// t.Run creates a subtest. This is great because:
		// 1. It gives a clear name to each test case in the output.
		// 2. If one subtest fails, others will still run.
		// 3. You can target specific subtests using `go test -run TestAdd/test_case_name`.
		t.Run(tc.name, func(t *testing.T) {
			// Execute the function we're testing with the inputs from the current test case.
			result := Add(tc.a, tc.b)

			// Check if the result matches the expected output.
			if result != tc.expected {
				// t.Errorf logs an error and marks the test as failed, but continues execution.
				// t.Fatalf logs an error, marks the test as failed, and stops execution of this specific subtest.
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

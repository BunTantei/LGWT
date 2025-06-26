package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Calculate positive number", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectCalculation(t, sum, expected)
	})

	t.Run("Calculate negative number", func(t *testing.T) {
		sum := Add(-2, -4)
		expected := -6
		assertCorrectCalculation(t, sum, expected)
	})

	t.Run("Calculate negative and postive number", func(t *testing.T) {
		sum := Add(-2, 4)
		expected := 2
		assertCorrectCalculation(t, sum, expected)
	})

	t.Run("Calculate postive and negative number", func(t *testing.T) {
		sum := Add(4, -2)
		expected := 2
		assertCorrectCalculation(t, sum, expected)
	})
}

func assertCorrectCalculation(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
